package publisher

import (
	"cloud.google.com/go/pubsub"
	"falabella.com/boss-message-library/core/domain/event"
	"falabella.com/boss-message-library/infra/secundary/config/properties"
)

var publisher = Publish

type HistoricalEventPublisher struct {
	topic *pubsub.Topic
}

const (
	EntityType = "history"
	EventType  = "historyUpdated"
)

func (p *HistoricalEventPublisher) Publish(historicalEvent *event.HistoricalEvent, country string, sourceSystem string, tenantId string) error {

	var attributes = make(map[string]string)
	attributes["country"] = country
	attributes["sourceSystem"] = sourceSystem
	attributes["tenantId"] = tenantId
	attributes["version"] = "1.0"
	CompleteHeader(attributes, EventType, EntityType)

	return publisher(p.topic, historicalEvent, attributes)

}

func NewHistoricalEventPublisher() *HistoricalEventPublisher {
	pid := properties.GetPublisherProperty().OrderStatus.ProjectId
	tid := properties.GetPublisherProperty().OrderStatus.TopicId
	return &HistoricalEventPublisher{
		topic: BuildTopic(pid, tid),
	}
}
