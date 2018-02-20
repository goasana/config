// Package config is an interface for dynamic configuration.
package config

import (
	"time"
)

// Config is an interface abstraction for dynamic configuration
type Config interface {
	// Config values
	Values
	// Config options
	Options() Options
	// Watch for changes
	Watch(path ...string) (Watcher, error)
	// Render config unusable
	Close() error
	// String name of config
	String() string
}

// Values is the interface for accessing config
// "path" could be a nested structure so it's composable
type Values interface {
	// Get cached value
	Get(path ...string) Value
	// Sets internal cached value
	Set(val interface{}, path ...string)
	// Deletes internal cached value
	Del(path ...string)
	// Returns vals as bytes
	Bytes() []byte
}

// Represent a value retrieved from the values loaded
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

// Source is the source from which config is loaded.
// This may be a file, a url, consul, env vars, etc.
type Source interface {
	// Loads ChangeSet from the source
	Read() (*ChangeSet, error)
	// Watch for source changes
	// Returns the entire changeset
	Watch() (SourceWatcher, error)
	// Name of source
	String() string
}

// Watcher is the config watcher
type Watcher interface {
	Next() (Value, error)
	Stop() error
}

// SourceWatcher allows you to watch a source for changes
// Next is a blocking call which returns the next
// ChangeSet update. Stop Renders the watcher unusable.
type SourceWatcher interface {
	Next() (*ChangeSet, error)
	Stop() error
}

// Reader takes a ChangeSet from a source and returns a single
// merged ChangeSet e.g reads ChangeSet as JSON and can merge down
type Reader interface {
	// Parse ChangeSets
	Parse(...*ChangeSet) (*ChangeSet, error)
	// As values
	Values(*ChangeSet) (Values, error)
	// Name of parser; json
	String() string
}

// ChangeSet represents a set an actual source
type ChangeSet struct {
	// The time at which the last change occured
	Timestamp time.Time
	// The raw data set for the change
	Data []byte
	// Hash of the source data
	Checksum string
	// The source of this change; file, consul, etcd
	Source string
}

type Options struct{}

type Option func(o *Options)

type SourceOptions struct{}

type SourceOption func(o *SourceOptions)
