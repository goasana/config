# Config [![GoDoc](https://godoc.org/github.com/micro/go-config?status.svg)](https://godoc.org/github.com/micro/go-config)

Go Config is a pluggable dynamic config library

Most configuration in application is statically configured via environment variables and config files. Hot reloading is usually left to the application or 
built in as complex logic into frameworks. Go-config separates out the concern of dynamic config into it's own library. 

## Config Format

Sources should return config in JSON format to operate with the default config reader

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

### Read File

```
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

### Go Vals

```go
// Get address. Set default to localhost as fallback
address := conf.Get("hosts", "database", "address").String("localhost")

// Get port. Set default to 3000 as fallback
port := conf.Get("hosts", "database", "port").Int(3000)
```
