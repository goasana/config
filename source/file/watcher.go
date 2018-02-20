package file

import (
	"errors"

	"github.com/micro/go-config/source"
	"gopkg.in/fsnotify.v1"
)

type Watcher struct {
	f *File

	fw   *fsnotify.Watcher
	exit chan bool
}

func newWatcher(f *File) (*Watcher, error) {
	fw, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	fw.Add(f.path)

	return &Watcher{
		f:    f,
		fw:   fw,
		exit: make(chan bool),
	}, nil
}

func (w *Watcher) Next() (*source.ChangeSet, error) {
	// is it closed?
	select {
	case <-w.exit:
		return nil, errors.New("watcher stopped")
	default:
	}

	// try get the event
	select {
	case <-w.fw.Events:
		c, err := w.f.Read()
		if err != nil {
			return nil, err
		}
		return c, nil
	case err := <-w.fw.Errors:
		return nil, err
	case <-w.exit:
		return nil, errors.New("watcher stopped")
	}
}

func (w *Watcher) Stop() error {
	return w.fw.Close()
}
