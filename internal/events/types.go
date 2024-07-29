package events

import "context"

// Handler is a function that handles an event
type Handler interface {
	Handle(ctx context.Context, message interface{}) error
}

// Subscriber is an interface that defines the methods for managing events
type Subscriber interface {
	Subscribe(topic string, handler Handler) error
}

// Publisher is an interface that defines the methods for publishing events
type Publisher interface {
	Publish(topicName string, payload []byte) error
}
