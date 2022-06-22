package event

import "github.com/google/uuid"

type HistoricalEvent struct {
	OrderId       string
	LogisticId    string
	CommerceId    string
	Event         string
	Description   string
	EventDateTime string
	EventId       uuid.UUID
	ExecutedBy    string
}
