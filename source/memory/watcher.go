package memory

import (
	"github.com/micro/go-config/source"
)

type Watcher struct {
	Id      string
	Updates chan *source.ChangeSet
	Source  *Source
}

func (w *Watcher) Next() (*source.ChangeSet, error) {
	cs := <-w.Updates
	return cs, nil
}

func (w *Watcher) Stop() error {
	w.Source.Lock()
	delete(w.Source.Watchers, w.Id)
	w.Source.Unlock()
	return nil
}
