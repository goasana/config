package consul

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/watch"
	"github.com/micro/go-config/source"
)

type watcher struct {
	name string

	wp   *watch.Plan
	ch   chan *source.ChangeSet
	exit chan bool
}

func newWatcher(key, addr, name string) (source.Watcher, error) {
	w := &watcher{
		name: name,
		ch:   make(chan *source.ChangeSet),
		exit: make(chan bool),
	}

	wp, err := watch.Parse(map[string]interface{}{"type": "keyprefix", "prefix": key})
	if err != nil {
		return nil, err
	}

	wp.Handler = w.handle

	if err := wp.Run(addr); err != nil {
		return nil, err
	}

	w.wp = wp

	return w, nil
}

func (w *watcher) handle(idx uint64, data interface{}) {
	if data == nil {
		return
	}

	kvs, ok := data.(api.KVPairs)
	if !ok {
		return
	}

	d := makeMap(kvs)

	b, err := json.Marshal(d)
	if err != nil {
		return
	}

	h := md5.New()
	h.Write(b)
	checksum := fmt.Sprintf("%x", h.Sum(nil))

	w.ch <- &source.ChangeSet{
		Source:   w.name,
		Data:     b,
		Checksum: checksum,
	}
}

func (w *watcher) Next() (*source.ChangeSet, error) {
	select {
	case cs := <-w.ch:
		return cs, nil
	case <-w.exit:
		return nil, errors.New("watcher stopped")
	}
}

func (w *watcher) Stop() error {
	select {
	case <-w.exit:
		return nil
	default:
		w.wp.Stop()
		close(w.exit)
	}
	return nil
}
