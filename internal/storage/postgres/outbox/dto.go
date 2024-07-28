package outbox

type OutboxEvent struct {
	ID        int
	EventName string
	Payload   []byte
}
