package instance

import (
	"falabella.com/boss-message-library/core/ports"
	"falabella.com/boss-message-library/infra/secundary/publisher"
)

var OrderStatusPublisher = func() ports.HistoricalEventPublisherPort {
	return publisher.NewHistoricalEventPublisher()
}
