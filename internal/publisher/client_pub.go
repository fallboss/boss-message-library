package publisher

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"falabella.com/boss-message-library/internal/config/log"
	"fmt"
	"sync"
)

var (
	ctx    = context.Background()
	once   sync.Once
	client *pubsub.Client
)

var logger = log.GetLogger()

func getClient(projectId string) *pubsub.Client {
	once.Do(func() {
		c, err := pubsub.NewClient(ctx, projectId)
		if err != nil {
			logger.Fatalf("Error creating client: %v", err)
		}
		client = c
	})
	return client
}

func BuildTopic(projectId string, topicId string) *pubsub.Topic {
	return getClient(projectId).Topic(topicId)
}

func Publish(t *pubsub.Topic, s interface{}, attributes map[string]string) error {
	b, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("error marshalling: %v", err)
	}

	message := pubsub.Message{Data: b, Attributes: attributes}

	logger.Infof("Message to publish: %s , Header: %v to", string(message.Data), message.Attributes)

	result := t.Publish(ctx, &message)

	_, err = result.Get(ctx)
	if err != nil {
		return fmt.Errorf("error publishing: %v", err)
	}
	return nil
}
