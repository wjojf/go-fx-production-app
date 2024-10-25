package storage

type OutboxProducer interface {
	StartProducing() error
	Stop()
}
