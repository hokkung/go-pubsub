package go_pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
)

// Publisher ...
type Publisher interface {
	Publish(ctx context.Context, req *PublishRequest) (*PublishResponse, error)
}

// BasePublisher manages base publisher
type BasePublisher struct {
	client *pubsub.Client
}

// NewPublisher creates an instance
func NewPublisher(client *pubsub.Client) *BasePublisher {
	return &BasePublisher{client: client}
}

// Publish publishes message to cloud pub/sub
func (p *BasePublisher) Publish(
	ctx context.Context,
	req *PublishRequest,
) (*PublishResponse, error) {
	topic := p.client.Topic(req.Topic)
	defer topic.Stop()

	result := topic.Publish(ctx, &pubsub.Message{Data: req.Data, Attributes: req.Attributes})
	serverID, err := result.Get(ctx)

	return &PublishResponse{
		ID: serverID,
	}, err
}
