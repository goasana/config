// Package config is an interface for dynamic configuration.
package config

import (
	"time"
)

// Config is an interface abstraction for dynamic configuration
type Config interface {
	Close() error
	Bytes() []byte
	Get(path ...string) Value
	Load(source ...Source) error
	Watch(path ...string) (Watcher, error)
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

// Source is the source from which config loaded
type Source interface {
	Read() (*ChangeSet, error)
	Watch() (SourceWatcher, error)
	String() string
}

// Watcher is the config watcher
type Watcher interface {
	Next() (Value, error)
	Stop() error
}

// SourceWatcher watches a source for changes
type SourceWatcher interface {
	Next() (*ChangeSet, error)
	Stop() error
}

// ChangeSet is a set of changes from a source
type ChangeSet struct {
	Data []byte
	Checksum string
	Updated time.Time
	Source string
}

type Options struct{}

type Option func(o *Options)

type SourceOptions struct{}

type SourceOption func(o *SourceOptions)
