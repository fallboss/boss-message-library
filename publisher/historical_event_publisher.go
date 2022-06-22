package publisher

import (
	"cloud.google.com/go/pubsub"
	"github.com/fallboss/boss-message-library/domain/event"
	"github.com/fallboss/boss-message-library/internal/config/log"
	publisher "github.com/fallboss/boss-message-library/internal/publisher"
)

var publish = publisher.Publish

type EventPublisher struct {
	topic *pubsub.Topic
}

var logger = log.GetLogger()

const (
	EntityType = "history"
	EventType  = "historyUpdated"
)

func NewEventPublisher() *EventPublisher {
	pid := GetPublisherProperty().HistoricalEvent.ProjectId
	tid := GetPublisherProperty().HistoricalEvent.TopicId

	logger.Infof("pid :", pid)
	logger.Infof("pid :", tid)
	return &EventPublisher{
		topic: publisher.BuildTopic(pid, tid),
	}
}

func (historicalEventPublisher *EventPublisher) Publish(e *event.HistoricalEvent, h map[string]string) error {

	var attributes = make(map[string]string)
	attributes["country"] = h["country"]
	attributes["sourceSystem"] = h["sourceSystem"]
	attributes["tenantId"] = h["tenantId"]
	attributes["version"] = "1.0"
	publisher.CompleteHeader(attributes, EventType, EntityType)

	return publish(historicalEventPublisher.topic, e, attributes)

}
