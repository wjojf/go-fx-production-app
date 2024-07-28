package outbox

type OutboxEvent struct {
	EventName string
	Payload   []byte
}
