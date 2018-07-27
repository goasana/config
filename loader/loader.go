// package loader manages loading from multiple sources
package loader

import (
	"github.com/micro/go-config/source"
)

// Loader manages loading sources
type Loader interface {
	// Load the sources
	Load(...source.Source) error
	// A Snapshot of loaded config
	Snapshot() (*Snapshot, error)
	// Force sync of sources
	Sync() error
	// Watch for changes
	Watch(...string) (Watcher, error)
	// Name of loader
	String() string
}

// Watcher lets you watch sources and returns a merged ChangeSet
type Watcher interface {
	Next() (*Snapshot, error)
	Stop()
}

// Snapshot is a merged ChangeSet
type Snapshot struct {
	// The merged ChangeSet
	ChangeSet *source.ChangeSet
	// Deterministic and comparable version of the snapshot
	Version string
}
