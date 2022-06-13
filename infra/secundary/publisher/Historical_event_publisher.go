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

func (p *HistoricalEventPublisher) Publish(historicalEvent *event.HistoricalEvent) error {

	var attributes = make(map[string]string)
	attributes["country"] = "CL"
	attributes["sourceSystem"] = "ANDES"
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

const (
	EntityType = "history"
	EventType  = "historyUpdated"
)
