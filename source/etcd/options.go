package etcd

import (
	"context"

	"github.com/micro/go-config/source"
)

type addressKey struct{}
type prefixKey struct{}
type stripPrefixKey struct{}
type usernameKey struct{}
type passwordKey struct{}

/*
type (
	addressKey     struct{}
	prefixKey      struct{}
	stripPrefixKey struct{}
	usernameKey    struct{}
	passwordKey    struct{}
)
*/

// WithAddress sets the consul address
func WithAddress(a string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, addressKey{}, a)
	}
}

// WithPrefix sets the key prefix to use
func WithPrefix(p string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, prefixKey{}, p)
	}
}

// StripPrefix indicates whether to remove the prefix from config entries, or leave it in place.
func StripPrefix(strip bool) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}

		o.Context = context.WithValue(o.Context, stripPrefixKey{}, strip)
	}
}

func WithUsername(name string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, userNameKey{}, name)
	}
}

func WithPassword(pass string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, passwordKey{}, pass)
	}
}
