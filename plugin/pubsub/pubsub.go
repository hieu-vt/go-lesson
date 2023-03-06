package pubsub

import (
	"context"
	"lesson-5-goland/pubsub"
)

type NatsPubSub interface {
	Publish(ctx context.Context, channel string, data *pubsub.Message) error
	Subscribe(ctx context.Context, channel string) (ch <-chan *pubsub.Message, close func())
}
