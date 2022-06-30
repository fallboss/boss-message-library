package publisher

import (
	"cloud.google.com/go/pubsub"
)

var publish = Publish

type EventPublisher struct {
	topic *pubsub.Topic
}

const (
	EntityType = "history"
	EventType  = "historyUpdated"
)

func NewEventPublisher() *EventPublisher {
	pid := GetPublisherProperty().HistoricalEvent.ProjectId
	tid := GetPublisherProperty().HistoricalEvent.TopicId

	return &EventPublisher{
		topic: BuildTopic(pid, tid),
	}
}

func (historicalEventPublisher *EventPublisher) Publish(e *HistoricalEvent, h map[string]string) error {

	var attributes = make(map[string]string)
	attributes["country"] = h["country"]
	attributes["sourceSystem"] = h["sourceSystem"]
	attributes["tenantId"] = h["tenantId"]
	attributes["version"] = "1.0"
	CompleteHeader(attributes, EventType, EntityType)

	return publish(historicalEventPublisher.topic, e, attributes)

}
