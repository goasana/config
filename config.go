// Package config is an interface for dynamic configuration.
package config

import (
	"context"

	"github.com/micro/go-config/source"
)

// Config is an interface abstraction for dynamic configuration
type Config interface {
	Close() error
	Bytes() []byte
	Get(path ...string) Value
	Load(source ...source.Source) error
	Watch(path ...string) (Watcher, error)
}

// Watcher is the config watcher
type Watcher interface {
	Next() (Value, error)
	Stop() error
}

// Value represents a value of any type
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

type Options struct {
	// for alternative data
	Context context.Context
}

type Option func(o *Options)
