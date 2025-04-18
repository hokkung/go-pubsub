package go_pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
)

// Subscriber ...
type Subscriber interface {
	Subscribe(ctx context.Context, subID string, handler func(context.Context, []byte, map[string]string) error) error
}

// SubscriberImpl ...
type SubscriberImpl struct {
	client *pubsub.Client
}

// NewSubscriber creates an instance
func NewSubscriber(client *pubsub.Client) *SubscriberImpl {
	return &SubscriberImpl{client: client}
}

// Subscribe subscribes message from cloud pub/sub
func (s *SubscriberImpl) Subscribe(
	ctx context.Context,
	subID string,
	handler func(context.Context, []byte, map[string]string) error,
) error {
	sub := s.client.Subscription(subID)
	return sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		err := handler(ctx, m.Data, m.Attributes)
		if err != nil {
			fmt.Println(err)
			m.Nack()
			return
		}

		m.Ack()
	})
}
