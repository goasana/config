# Config [![GoDoc](https://godoc.org/github.com/micro/go-config?status.svg)](https://godoc.org/github.com/micro/go-config)

Go Config is a pluggable dynamic config library

Most configuration in application is statically configured via environment variables and config files. Hot reloading is usually left to the application or 
built in as complex logic into frameworks. Go-config separates out the concern of dynamic config into it's own library. 

## Features

- Dynamic config
- Pluggable sources
- Source merging
- Default values
- Config watcher

## Sources

The following sources for config are supported

- [consul](https://godoc.org/github.com/micro/go-config/source/consul) - read from consul
- [envvar](https://godoc.org/github.com/micro/go-config/source/envvar) - read from environment variables
- [file](https://godoc.org/github.com/micro/go-config/source/file) - read from file
- [flag](https://godoc.org/github.com/micro/go-config/source/flag) - read from flags
- [memory](https://godoc.org/github.com/micro/go-config/source/memory) - read from memory

## Interface

The interface is very simple. It supports multiple config sources, watching and default fallback values.

```go
type Config interface {
        Close() error
        Bytes() []byte
        Get(path ...string) Value
        Load(source ...source.Source) error
        Watch(path ...string) (Watcher, error)
}
```

## Source

A [Source](https://godoc.org/github.com/micro/go-config/source#Source) is the source of config. 

It can be env vars, a file, a key value store. Anything which conforms to the Source interface.

### Interface

```go
// Source is the source from which config is loaded
type Source interface {
	Read() (*ChangeSet, error)
	Watch() (Watcher, error)
	String() string
}

// ChangeSet represents a set of changes from a source
type ChangeSet struct {
	Data      []byte
	Checksum  string
	Timestamp time.Time
	Source    string
}
```

### Format

Sources should return config in JSON format to operate with the default config reader

The [Reader](https://godoc.org/github.com/micro/go-config/reader#Reader) defaults to json but can be swapped out to any other format.

```
{
	"path": {
		"to": {
			"key": ["foo", "bar"]
		}
	}
}
```

## Usage

Assuming the following config file

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

```
import "github.com/micro/go-config/source/file"

// Create new config
conf := config.NewConfig()

// Load file source
conf.Load(file.NewSource(
	file.WithPath("/tmp/config.json"),
))
```

### Scan

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

### Go Vals

```go
// Get address. Set default to localhost as fallback
address := conf.Get("hosts", "database", "address").String("localhost")

// Get port. Set default to 3000 as fallback
port := conf.Get("hosts", "database", "port").Int(3000)
```

### Watch

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

