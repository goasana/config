package vault

import (
	"errors"
	"github.com/goasana/config/source"

	"github.com/hashicorp/vault/api"
)

type watcher struct {
	c    *api.Client
	exit chan bool
}

func newWatcher(c *api.Client) *watcher {
	return &watcher{
		c:    c,
		exit: make(chan bool),
	}
}

func (w *watcher) Next() (*source.ChangeSet, error) {
	<-w.exit
	return nil, errors.New("url watcher stopped")
}

func (w *watcher) Stop() error {
	select {
	case <-w.exit:
	default:
	}
	return nil
}
