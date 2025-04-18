package go_pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
)

// NewClient creates an instance
func NewClient(ctx context.Context, cfg *Config) (*pubsub.Client, error) {
	return pubsub.NewClient(ctx, cfg.ProjectID)
}

// ProvideClient provides client
func ProvideClient(ctx context.Context, cfg *Config) (pubsub.Client, error) {
	c, err := NewClient(ctx, cfg)
	return *c, err
}
