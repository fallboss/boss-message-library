package publisher

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

var (
	ctx        = context.Background()
	clientOnce sync.Once
	client     *pubsub.Client
)

var logg = GetLogger()

func getClient(projectId string) *pubsub.Client {
	clientOnce.Do(func() {
		c, err := pubsub.NewClient(ctx, projectId)
		if err != nil {
			logg.Fatalf("Error creating client: %v", err)
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

	logg.Infof("Header to publish: %v", message.Attributes)
	logg.Infof("Message to publish : %v", string(message.Data))

	result := t.Publish(ctx, &message)

	_, err = result.Get(ctx)
	if err != nil {
		return fmt.Errorf("error publishing: %v", err)
	}
	return nil
}
