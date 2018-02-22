# microcli Source

The microcli source reads config from parsed flags via a cli.Context.

## Format

We expect the use of the `micro/cli` package. Upper case flags will be lower cased. Dashes will be used as delimiters for nesting.

### Example

```go
dbAddress := flag.String("database-address", "127.0.0.1", "the db address")
dbPort := flag.Int("database-port", 3306, "the db port)
```

Becomes

```json
{
    "database": {
        "address": "127.0.0.1",
        "port": 3306
    }
}
```

## New and Load Source

Because a cli.Context is needed to retrieve the flags and their values, it is recommended to build your source from within a cli.Action.

```go

func main() {

    // New Service
    service := micro.NewService(
        micro.Name("example"),
        micro.Flags([]cli.Flag{
            cli.StringFlag{Name: "database-name", Usage: "database name"},
        }),
    )

    var clisrc source.Source
    service.Init(
        micro.Action(func(c *cli.Context) {
            clisrc = microcli.NewSource(c)
            // Alternatively, just setup your config right here
        }),
    )
    
    // ... Load and use that source ...
    conf := config.NewConfig()
    conf.Load(clisrc)
}
```
