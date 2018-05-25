// Package configmap config is an interface for dynamic configuration.
package configmap

import (
	"crypto/md5"
	"encoding/json"
	"fmt"

	"github.com/micro/go-config/source"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type configmap struct {
	opts       source.Options
	client     *kubernetes.Clientset
	name       string
	namespace  string
	configPath string
}

var (
	DefaultName       = "micro"
	DefaultConfigPath = ""
	DefaultNamespace  = "default"
)

func (k *configmap) Read() (*source.ChangeSet, error) {
	cmp, err := k.client.CoreV1().ConfigMaps(k.namespace).Get(k.name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	data := makeMap(cmp.Data)

	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error reading source: %v", err)
	}

	h := md5.New()
	h.Write(b)
	checksum := fmt.Sprintf("%x", h.Sum(nil))

	return &source.ChangeSet{
		Source:   k.String(),
		Data:     b,
		Checksum: checksum,
	}, nil
}

func (k *configmap) String() string {
	return "configmap"
}

func (k *configmap) Watch() (source.Watcher, error) {
	w, err := newWatcher(k.name, k.namespace, k.client)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func NewSource(opts ...source.Option) source.Source {
	var (
		options    source.Options
		name       = DefaultName
		configPath = DefaultConfigPath
		namespace  = DefaultNamespace
	)

	for _, o := range opts {
		o(&options)
	}

	if options.Context != nil {
		prefix, ok := options.Context.Value(prefixKey{}).(string)
		if ok {
			name = prefix
		}

		cfg, ok := options.Context.Value(configPathKey{}).(string)
		if ok {
			configPath = cfg
		}

		sname, ok := options.Context.Value(nameKey{}).(string)
		if ok {
			name = sname
		}

		ns, ok := options.Context.Value(namespaceKey{}).(string)
		if ok {
			namespace = ns
		}
	}

	// TODO handle if the client fails what to do current return does not support error
	client, _ := getClient(configPath)
	// if err != nil {
	// 	fmt.Errorf("unable to get the configmap client: %v", err)
	// }

	return &configmap{
		client:     client,
		opts:       options,
		name:       name,
		configPath: configPath,
		namespace:  namespace,
	}
}
