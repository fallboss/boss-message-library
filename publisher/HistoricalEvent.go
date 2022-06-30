package publisher

import "github.com/google/uuid"

type HistoricalEvent struct {
	OrderId       string    `json:"orderId,omitempty"`
	LogisticId    string    `json:"logisticId,omitempty"`
	CommerceId    string    `json:"commerceId,omitempty"`
	Event         string    `json:"event,omitempty"`
	Description   string    `json:"description,omitempty"`
	EventDateTime string    `json:"eventDateTime,omitempty"`
	EventId       uuid.UUID `json:"eventId,omitempty"`
	ExecutedBy    string    `json:"executeBy,omitempty"`
}
