package falabella_boss_message_go_library

import (
	"cloud.google.com/go/pubsub"
	"falabella.com/boss-message-library/domain/event"
	"falabella.com/boss-message-library/internal/config/properties"
	"falabella.com/boss-message-library/internal/publisher"
)

var publish = publisher.Publish

type eventPublisher struct {
	topic *pubsub.Topic
}

const (
	EntityType = "history"
	EventType  = "historyUpdated"
)

func (historicalEventPublisher *eventPublisher) Publish(e *event.HistoricalEvent, h map[string]string) error {

	var attributes = make(map[string]string)
	attributes["country"] = h["country"]
	attributes["sourceSystem"] = h["sourceSystem"]
	attributes["tenantId"] = h["tenantId"]
	attributes["version"] = "1.0"
	publisher.CompleteHeader(attributes, EventType, EntityType)

	return publish(historicalEventPublisher.topic, e, attributes)

}

func NewEventPublisher() *eventPublisher {
	pid := properties.GetPublisherProperty().HistoricalEvent.ProjectId
	tid := properties.GetPublisherProperty().HistoricalEvent.TopicId
	return &eventPublisher{
		topic: publisher.BuildTopic(pid, tid),
	}
}
