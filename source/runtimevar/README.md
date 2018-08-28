# Runtimevar Source

The runtimevar source is a source for the [Go Cloud](https://github.com/google/go-cloud) runtimevar package

This package takes a [driver.Watcher](https://godoc.org/github.com/google/go-cloud/runtimevar/driver#Watcher) 
and then allows you to use it as a backend source. We expect the 
[Snapshot](https://godoc.org/github.com/google/go-cloud/runtimevar#Snapshot) value to be `[]byte`. 
We use the built in encoder to decode the value. This defaults to json.

## Example

A runtimevar config format in json

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

## New Source

Specify runtimevar source with watcher. It will panic if not specified.

```go
srv := runtimevar.NewSource(
	runtimevar.WithDriver(dv),
)
```

## Config Format

To load different runtimevar formats e.g yaml, toml, xml you must specify an encoder

```
e := toml.NewEncoder()

src := runtimevar.NewSource(
        runtimevar.WithDriver(dv),
	source.WithEncoder(e),
)
```

## Load Source

Load the source into config

```go
// Create new config
conf := config.NewConfig()

// Load runtimevar source
conf.Load(src)
```

