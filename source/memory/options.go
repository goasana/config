package memory

import (
	"github.com/micro/go-config/source"

	"context"
)

type dataKey struct{}

// WithData allows the source data to be set
func WithData(d []byte) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, dataKey{}, d)
	}
}
