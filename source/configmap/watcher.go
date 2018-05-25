package configmap

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/micro/go-config/source"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type watcher struct {
	name      string
	namespace string
	client    *kubernetes.Clientset
	st        cache.Store
	ct        cache.Controller
	ch        chan *source.ChangeSet
	exit      chan bool
	stop      chan struct{}
}

func newWatcher(n, ns string, c *kubernetes.Clientset) (source.Watcher, error) {
	w := &watcher{
		name:      n,
		namespace: ns,
		client:    c,
		ch:        make(chan *source.ChangeSet),
		exit:      make(chan bool),
		stop:      make(chan struct{}),
	}

	lw := cache.NewListWatchFromClient(w.client.CoreV1().RESTClient(), "ConfigMap", w.namespace, fields.OneTermEqualSelector("metadata.name", w.name))
	st, ct := cache.NewInformer(
		lw,
		&v12.ConfigMap{},
		time.Second*30,
		cache.ResourceEventHandlerFuncs{
			UpdateFunc: w.handle,
		},
	)

	go ct.Run(w.stop)

	w.ct = ct
	w.st = st

	return w, nil
}

func (w *watcher) handle(oldCmp interface{}, newCmp interface{}) {
	if newCmp == nil {
		return
	}

	data := makeMap(newCmp.(v12.ConfigMap).Data)

	b, err := json.Marshal(data)
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

// Next
func (w *watcher) Next() (*source.ChangeSet, error) {
	select {
	case cs := <-w.ch:
		return cs, nil
	case <-w.exit:
		return nil, errors.New("watcher stopped")
	}
}

// Stop
func (w *watcher) Stop() error {
	select {
	case <-w.exit:
		return nil
	case <-w.stop:
		return nil
	default:
		close(w.exit)
	}
	return nil
}
