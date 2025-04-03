package graph

import "context"

type contextKey uint32

const (
	optimizationsEnabled contextKey = iota
)

func ContextWithOptimizations(ctx context.Context, enabled bool) context.Context {
	return context.WithValue(ctx, optimizationsEnabled, enabled)
}

func OptimizationsEnabledFromContext(ctx context.Context) *bool {
	enabled := ctx.Value(optimizationsEnabled)
	if enabled != nil {
		enabledBool, ok := enabled.(bool)
		if ok {
			return &enabledBool
		}
	}
	return nil
}
