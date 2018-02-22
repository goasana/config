package envvar

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/imdario/mergo"
	"github.com/micro/go-config/source"
)

type envvar struct {
	prefix string
	opts   source.Options
}

func (e *envvar) Read() (*source.ChangeSet, error) {
	var changes map[string]interface{}

	for _, env := range os.Environ() {
		if len(e.prefix) > 0 {
			if !strings.HasPrefix(env, e.prefix) {
				continue
			}

			env = strings.TrimPrefix(env, e.prefix)
		}

		env = strings.ToLower(env)
		pair := strings.Split(env, "=")
		value := pair[1]
		keys := strings.Split(pair[0], "_")
		reverse(keys)

		tmp := make(map[string]interface{})
		for i, k := range keys {
			if i == 0 {
				tmp[k] = value
				continue
			}

			tmp = map[string]interface{}{k: tmp}
		}

		if err := mergo.Map(&changes, tmp); err != nil {
			return nil, err
		}
	}

	b, err := json.Marshal(changes)
	if err != nil {
		return nil, err
	}

	h := md5.New()
	h.Write(b)
	checksum := fmt.Sprintf("%x", h.Sum(nil))

	return &source.ChangeSet{
		Data:      b,
		Checksum:  checksum,
		Timestamp: time.Now(),
		Source:    e.String(),
	}, nil
}

func reverse(ss []string) {
	for i := len(ss)/2 - 1; i >= 0; i-- {
		opp := len(ss) - 1 - i
		ss[i], ss[opp] = ss[opp], ss[i]
	}
}

func (e *envvar) Watch() (source.Watcher, error) {
	return newWatcher()
}

func (e *envvar) String() string {
	return "envvar"
}

// NewSource returns a config source for parsing ENV variables.
// Underscores are delimiters for nesting, and all keys are lowercased.
//
// Example:
//      "DATABASE_SERVER_HOST=localhost" will convert to
//
//      {
//          "database": {
//              "server": {
//                  "host": "localhost"
//              }
//          }
//      }
func NewSource(opts ...source.Option) source.Source {
	var options source.Options
	for _, o := range opts {
		o(&options)
	}

	var prefix string
	if options.Context != nil {
		if p, ok := options.Context.Value(prefixKey{}).(string); ok {
			prefix = p
		}
	}
	return &envvar{prefix: prefix, opts: options}
}
