# Config [![GoDoc](https://godoc.org/github.com/micro/go-config?status.svg)](https://godoc.org/github.com/micro/go-config)

Go Config is a pluggable dynamic config library

Most config in applications are statically configured or include complex logic to load from multiple sources. 
Go-config makes this easy, pluggable and mergeable. You'll never have to deal with config in the same way again.

## Features

- **Dynamic** - load config on the fly as you need it
- **Pluggable** - choose which source to load from; file, envvar, consul
- **Mergeable** - merge and override multiple config sources
- **Fallback** - specify fallback values where keys don't exist
- **Watch** - Watch the config for changes

## Getting Started

- [Sources](#sources)
- [Formats](#formats)
- [Config](#config)
- [Usage](#usage)
- [FAQ](#faq)

## Sources

Sources are backends from which config is loaded. The following sources for config are supported.

- [configmap](https://github.com/micro/go-config/tree/master/source/configmap) - read from k8s configmap
- [consul](https://github.com/micro/go-config/tree/master/source/consul) - read from consul
- [etcd](https://github.com/micro/go-config/tree/master/source/etcd) - read from etcd v3
- [envvar](https://github.com/micro/go-config/tree/master/source/envvar) - read from environment variables
- [file](https://github.com/micro/go-config/tree/master/source/file) - read from file
- [flag](https://github.com/micro/go-config/tree/master/source/flag) - read from flags
- [grpc](https://github.com/micro/go-config/tree/master/source/grpc) - read from grpc server
- [memory](https://github.com/micro/go-config/tree/master/source/memory) - read from memory
- [microcli](https://github.com/micro/go-config/tree/master/source/microcli) - read from micro cli flags

TODO:

- vault
- git url

## Formats

Encoders handle source encoding formats. The following encoding formats are supported:

- json
- yaml
- toml
- xml
- hcl 

Default encoder is json with format:

```json
{
	"path": {
		"to": {
			"key": ["foo", "bar"]
		}
	}
}
```

## Config 

Top level config is an interface. It supports multiple sources, watching and fallback values.

### Interface

```go

// Config is an interface abstraction for dynamic configuration
type Config interface {
	// Stop the config loader/watcher
	Close() error
	// Get the whole config as raw output
	Bytes() []byte
	// Force a source changeset sync
	Sync() error
	// Get a value from the config
	Get(path ...string) reader.Value
	// Load config sources
	Load(source ...source.Source) error
	// Watch a value for changes
	Watch(path ...string) (Watcher, error)
}
```

### Value

The `config.Get` method returns a `reader.Value` which can cast to any type with a fallback value

```go
type Value interface {
	Bool(def bool) bool
	Int(def int) int
	String(def string) string
	Float64(def float64) float64
	Duration(def time.Duration) time.Duration
	StringSlice(def []string) []string
	StringMap(def map[string]string) map[string]string
	Scan(val interface{}) error
	Bytes() []byte
}
```

## Usage

- [Sample Config](#sample-config)
- [Load File](#load-file)
- [Scan Value](#scan-value)
- [Cast Value](#cast-value)
- [Watch Path](#watch-path)
- [Merge Sources](#merge-sources)
- [Set Source Encoder](#set-source-encoder)
- [Add Reader Encoder](#add-reader-encoder)

### Sample Config

```json
{
    "hosts": {
        "database": {
            "address": "10.0.0.1",
            "port": 3306
        },
        "cache": {
            "address": "10.0.0.2",
            "port": 6379
        }
    }
}
```

### Load File

```go
import "github.com/micro/go-config/source/file"

// Create new config
conf := config.NewConfig()

// Load file source
conf.Load(file.NewSource(
	file.WithPath("/tmp/config.json"),
))
```

### Scan Value

```go
type Host struct {
	Address string `json:"address"`
	Port int `json:"port"`
}

var host Host

conf.Get("hosts", "database").Scan(&host)

// 10.0.0.1 3306
fmt.Println(host.Address, host.Port)
```

### Cast Value

```go
// Get address. Set default to localhost as fallback
address := conf.Get("hosts", "database", "address").String("localhost")

// Get port. Set default to 3000 as fallback
port := conf.Get("hosts", "database", "port").Int(3000)
```

### Watch Path

Watch a path for changes. When the file changes the new value will be made available.

```go
w, err := conf.Watch("hosts", "database")
if err != nil {
	// do something
}

// wait for next value
v, err := w.Next()
if err != nil {
	// do something
}

var host Host

v.Scan(&host)
```

### Merge Sources

Multiple sources can be loaded and merged. Merging priority is in reverse order. 

```go
conf := config.NewConfig()


conf.Load(
	// base config from env
	envvar.NewSource(),
	// override env with flags
	flag.NewSource(),
	// override flags with file
	file.NewSource(
		file.WithPath("/tmp/config.json"),
	),
)
```

### Set Source Encoder

The default encoder is json.

```go
e := yaml.NewEncoder()

s := consul.NewSource(
	source.WithEncoder(e),
)
```

### Add Reader Encoder

The reader supports multiple encoders.

Add a new encoder to a reader like so:

```go
e := yaml.NewEncoder()

r := json.NewReader(
	reader.WithEncoder(e),
)
```

## FAQ

### How is this different from Viper?

[Viper](https://github.com/spf13/viper) and go-config are solving the same problem. Go-config provides a different interface and is part of the larger micro 
ecosystem of tooling.

