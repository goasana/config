package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	cetcd "github.com/coreos/etcd/clientv3"
	"github.com/micro/go-config/source"
)

// Currently a single etcd reader
type etcd struct {
	prefix      string
	stripPrefix string
	opts        source.Options
	client      *cetcd.Client
	cerr        error
}

var (
	DefaultPrefix = "/micro/config/"
)

func (c *etcd) Read() (*source.ChangeSet, error) {
	if c.cerr != nil {
		return nil, c.cerr
	}

	rsp, err := c.client.Get(context.Background(), c.prefix, cetcd.WithPrefix())
	if err != nil {
		return nil, err
	}

	if rsp == nil || len(rsp.Kvs) == 0 {
		return nil, fmt.Errorf("source not found: %s", c.prefix)
	}

	data := makeMap(rsp.Kvs, c.stripPrefix)

	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error reading source: %v", err)
	}

	cs := &source.ChangeSet{
		Timestamp: time.Now(),
		Source:    c.String(),
		Data:      b,
	}
	cs.Checksum = cs.Sum()

	return cs, nil
}

func (c *etcd) String() string {
	return "etcd"
}

func (c *etcd) Watch() (source.Watcher, error) {
	if c.cerr != nil {
		return nil, c.cerr
	}
	cs, err := c.Read()
	if err != nil {
		return nil, err
	}
	return newWatcher(c.prefix, c.stripPrefix, c.client.Watcher, cs)
}

func NewSource(opts ...source.Option) source.Source {
	var options source.Options

	for _, o := range opts {
		o(&options)
	}

	endpoints := []string{"localhost:2379"}

	// check if there are any addrs
	if options.Context != nil {
		a, ok := options.Context.Value(addressKey{}).(string)
		if ok {
			addr, port, err := net.SplitHostPort(a)
			if ae, ok := err.(*net.AddrError); ok && ae.Err == "missing port in address" {
				port = "2379"
				addr = a
				endpoints = []string{fmt.Sprintf("%s:%s", addr, port)}
			} else if err == nil {
				endpoints = []string{fmt.Sprintf("%s:%s", addr, port)}
			}
		}
	}

	// use default config
	client, err := cetcd.New(cetcd.Config{
		Endpoints: endpoints,
	})

	prefix := DefaultPrefix
	sp := ""
	if options.Context != nil {
		f, ok := options.Context.Value(prefixKey{}).(string)
		if ok {
			prefix = f
		}

		if b, ok := options.Context.Value(stripPrefixKey{}).(bool); ok && b {
			sp = prefix
		}
	}

	return &etcd{
		prefix:      prefix,
		stripPrefix: sp,
		opts:        options,
		client:      client,
		cerr:        err,
	}
}
