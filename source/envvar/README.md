# Env Var Source

The envvar source reads config from environment variables

## Format

We expect environment variables to be in the standard format of FOO=bar

Keys are converted to lowercase and split on underscore.


### Example

```
DATABASE_ADDRESS=127.0.0.1
DATABASE_PORT=3306
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

### Namespace

Env vars can be namespaced so we only have access to a subset

```
MICRO_DATABASE_ADDRESS=127.0.0.1
MICRO_DATABASE_PORT=3306
```

The prefix will be trimmed on access

## New Source

Specify source with data

```go
envvarSource := envvar.NewSource(
	// optionally specify prefix
	envvar.WithPrefix("MICRO"),
)
```

## Load Source

Load the source into config

```go
// Create new config
conf := config.NewConfig()

// Load file source
conf.Load(envvarSource)
```
