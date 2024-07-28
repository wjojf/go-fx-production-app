package events

type UserCreatedPayload struct {
	EventID string `json:"event_id"`
	UserID  string `json:"user_id"`
}
