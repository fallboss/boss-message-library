package ports

import "falabella.com/boss-message-library/core/domain/event"

type HistoricalEventPublisherPort interface {
	Publish(historicalEvent *event.HistoricalEvent) error
}
