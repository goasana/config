# Config [![GoDoc](https://godoc.org/github.com/micro/go-config?status.svg)](https://godoc.org/github.com/micro/go-config)

Go Config is a pluggable dynamic config library

Most configuration in application is statically configured via environment variables and config files. Hot reloading is usually left to the application or 
built in as complex logic into frameworks. Go-config separates out the concern of dynamic config into it's own library. 

## Config Format

Sources should return config in JSON format to operate with the default config interface

```
{
	"path": {
		"to": {
			"key": ["foo", "bar"]
		}
	}
}
```
