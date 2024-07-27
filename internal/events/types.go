package events

import "context"

// Topic is a type that represents the topic of an event (example "user.created"(string) or *kafka.TopicPartition{Topic: "user.created"})
type Topic any

// Handler is a function that handles an event
type Handler interface {
	Handle(ctx context.Context, payload any)
}

// Manager is an interface that defines the methods for managing events
type Subscriber interface {
	Subscribe(topic Topic, handler Handler) error
}
