// Package config is an interface for dynamic configuration.
package config

import (
	"time"
)

// Config is an interface abstraction for dynamic configuration
type Config interface {
	Close() error
	Load(source ...Source) error
	Watch(path ...string) (Watcher, error)
	Values
}

// Values is the interface for accessing config
// "path" could be a nested structure so it's composable
type Values interface {
	Get(path ...string) Value
	Bytes() []byte
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
	// Loads ChangeSet from the source
	Read() (*ChangeSet, error)
	// Watch for source changes
	// Returns the entire changeset
	Watch() (SourceWatcher, error)
	// Name of source; env, file, consul
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
	// The time at which the last change occured
	Timestamp time.Time
	// The raw data set for the change
	Data []byte
	// Hash of the source data
	Checksum string
	// The source of this change
	Source string
}

type Options struct{}

type Option func(o *Options)

type SourceOptions struct{}

type SourceOption func(o *SourceOptions)
