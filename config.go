// Package config is an interface for dynamic configuration.
package config

import (
	"context"

	"github.com/micro/go-config/reader"
	"github.com/micro/go-config/source"
)

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

// Watcher is the config watcher
type Watcher interface {
	Next() (reader.Value, error)
	Stop() error
}

type Options struct {
	Reader reader.Reader
	Source []source.Source

	// for alternative data
	Context context.Context
}

type Option func(o *Options)

// NewConfig returns new config
func NewConfig(opts ...Option) Config {
	return newConfig(opts...)
}
