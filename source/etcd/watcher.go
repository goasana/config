package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	cetcd "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/micro/go-config/source"
)

type watcher struct {
	name        string
	stripPrefix string

	ch   chan *source.ChangeSet
	exit chan bool
}

func newWatcher(key, name, stripPrefix string, wc cetcd.Watcher) (source.Watcher, error) {
	w := &watcher{
		name:        name,
		stripPrefix: stripPrefix,
		ch:          make(chan *source.ChangeSet),
		exit:        make(chan bool),
	}

	ch := wc.Watch(context.Background(), key, cetcd.WithPrefix())

	go w.run(wc, ch)

	return w, nil
}

func (w *watcher) handle(evs []*cetcd.Event) {
	var kvs []*mvccpb.KeyValue

	for _, v := range evs {
		kvs = append(kvs, v.Kv)
	}

	d := makeMap(kvs, w.stripPrefix)

	b, err := json.Marshal(d)
	if err != nil {
		return
	}
	cs := &source.ChangeSet{
		Timestamp: time.Now(),
		Source:    w.name,
		Data:      b,
		Format:    "json",
	}
	cs.Checksum = cs.Sum()
	w.ch <- cs
}

func (w *watcher) run(wc cetcd.Watcher, ch cetcd.WatchChan) {
	for {
		select {
		case rsp, ok := <-ch:
			if !ok {
				return
			}
			w.handle(rsp.Events)
		case <-w.exit:
			wc.Close()
			return
		}
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
		close(w.exit)
	}
	return nil
}
