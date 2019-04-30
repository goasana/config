# Kubernetes ConfigMap Source (configmap)

The configmap source reads config from a kubernetes configmap key/values

## Kubernetes ConfigMap Format

The configmap source expects keys under a namespace default to `default` and a confimap default to `asana`

```shell
// we recommend to setup your variables from multiples files example:
$ kubectl create configmap asana --namespace default --from-file=./testdata

// verify if were set correctly with
$ kubectl get configmap asana --namespace default
{
    "apiVersion": "v1",
    "data": {
        "config": "host=0.0.0.0\nport=1337",
        "mongodb": "host=127.0.0.1\nport=27017\nuser=user\npassword=password",
        "redis": "url=redis://127.0.0.1:6379/db01"
    },
    "kind": "ConfigMap",
    "metadata": {
        ...
        "name": "asana",
        "namespace": "default",
        ...
    }
}
```

Keys are split on `\n` and `=` this is because the way kubernetes saves the data is `map[string][string]`.

```go
// the example above "mongodb": "host=127.0.0.1\nport=27017\nuser=user\npassword=password" will be accessible as:
conf.Get("mongodb", "host") // 127.0.0.1
conf.Get("mongodb", "port") // 27017
```

## New Source

Specify source with data

```go
configmapSource := configmap.NewSource(
	// optionally specify a namespace; default to default
	configmap.WithNamespace("kube-public"),
	// optionally specify name for ConfigMap; defaults asana
	configmap.WithName("asana-config"),
    // optionally strip the provided path to a kube config file mostly used outside of a cluster, defaults to "" for in cluster support.
    configmap.WithConfigPath($HOME/.kube/config),
)
```

## Load Source

Load the source into config

```go
// Create new config
conf := config.NewConfig()

// Load file source
conf.Load(configmapSource)
```

## Running Go Tests

### Requirements

Have a kubernetes cluster running (external or minikube) have a valid `kubeconfig` file.

```shell
// Setup testing configmaps feel free to remove them after testing.
$ cd source/configmap
$ kubectl create configmap asana --from-file=./testdata
$ kubectl create configmap asana --from-file=./testdata --namespace kube-public
$ kubectl create configmap asana-config --from-file=./testdata
$ kubectl create configmap asana-config --from-file=./testdata --namespace kube-public
$ go test -v -cover
```

```shell
// To clean up the testing configmaps
$ kubectl delete configmap asana --all-namespaces
$ kubectl delete configmap asana-config --all-namespaces
```

## Todos
- [ ] add more test cases including watchers
- [ ] add support for prefixing either using namespace or a custom `string` passed as `WithPrefix`
- [ ] a better way to test without manual setup from the user.
- [ ] add test examples.
- [ ] open to suggestions and feedback please let me know what else should I add.

**stay tuned for kubernetes secret support as an source.**