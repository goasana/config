package envvar

import (
	"context"

	"github.com/micro/go-config/source"
)

type prefixKey struct{}

// WithPrefix sets the environment variable prefix to scope to.
// This prefix will be removed from the actual config entries.
func WithPrefix(p string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, prefixKey{}, p)
	}
}
