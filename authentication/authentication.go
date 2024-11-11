package authentication

import (
	"context"
	"slices"
)

type Privilege string

type privilegeKey struct{}

func Grant(ctx context.Context, ps ...Privilege) context.Context {
	if ctx.Value(privilegeKey{}) != nil {
		panic("Grant called multiple times")
	}
	return context.WithValue(ctx, privilegeKey{}, ps)
}

type checkedKey struct{}

func Check(ctx context.Context, ps ...Privilege) (_ context.Context, ok bool) {
	granted, ok := ctx.Value(privilegeKey{}).([]Privilege)
	if !ok {
		return ctx, false
	}

	for _, p := range ps {
		if !slices.Contains(granted, p) {
			return ctx, false
		}

	}
	return context.WithValue(ctx, checkedKey{}, struct{}{}), true
}

func Must(ctx context.Context) bool {
	return ctx.Value(checkedKey{}) != nil
}
