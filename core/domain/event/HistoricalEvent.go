package event

import "github.com/google/uuid"

type HistoricalEvent struct {
	id            uuid.UUID
	logisticId    string
	commerceId    string
	event         string
	description   string
	eventDateTime string
	eventId       uuid.UUID
	executedBy    string
}
