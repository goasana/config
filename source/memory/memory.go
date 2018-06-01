// Package memory is a memory source
package memory

import (
	"crypto/md5"
	"fmt"
	"sync"
	"time"

	"github.com/micro/go-config/source"
	"github.com/pborman/uuid"
)

type memory struct {
	sync.RWMutex
	ChangeSet *source.ChangeSet
	Watchers  map[string]*watcher
}

func (s *memory) Read() (*source.ChangeSet, error) {
	s.RLock()
	cs := &source.ChangeSet{
		Timestamp: s.ChangeSet.Timestamp,
		Data:      s.ChangeSet.Data,
		Checksum:  s.ChangeSet.Checksum,
		Source:    s.ChangeSet.Source,
	}
	s.RUnlock()
	return cs, nil
}

func (s *memory) Watch() (source.Watcher, error) {
	w := &watcher{
		Id:      uuid.NewUUID().String(),
		Updates: make(chan *source.ChangeSet, 100),
		Source:  s,
	}

	s.Lock()
	s.Watchers[w.Id] = w
	s.Unlock()
	return w, nil
}

// Update allows manual updates of the config data.
func (s *memory) Update(c *source.ChangeSet) {
	// don't process nil
	if c == nil {
		return
	}

	// hash the file
	h := md5.New()
	h.Write(c.Data)
	checksum := fmt.Sprintf("%x", h.Sum(nil))

	s.Lock()
	// update changeset
	s.ChangeSet = &source.ChangeSet{
		Checksum:  checksum,
		Data:      c.Data,
		Format:    c.Format,
		Source:    "memory",
		Timestamp: time.Now(),
	}

	// update watchers
	for _, w := range s.Watchers {
		select {
		case w.Updates <- s.ChangeSet:
		default:
		}
	}
	s.Unlock()
}

func (s *memory) String() string {
	return "memory"
}

func NewSource(opts ...source.Option) source.Source {
	var options source.Options
	for _, o := range opts {
		o(&options)
	}

	s := &memory{
		Watchers: make(map[string]*watcher),
	}

	if options.Context != nil {
		c, ok := options.Context.Value(changeSetKey{}).(*source.ChangeSet)
		if ok {
			s.Update(c)
		}
	}

	return s
}
