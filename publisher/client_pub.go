package publisher

import (
	"bytes"
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

var (
	ctx        = context.Background()
	clientOnce sync.Once
	client     *pubsub.Client
)

func getClient(projectId string) *pubsub.Client {
	clientOnce.Do(func() {
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

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, message.Data, "", "\t")
	if err != nil {
		logger.Errorf("JSON parse error: ", err)
	}
	log.Println("Header to publish: ", message.Attributes)
	log.Println("Message to publish :", string(prettyJSON.Bytes()), message.Attributes)

	result := t.Publish(ctx, &message)

	_, err = result.Get(ctx)
	if err != nil {
		return fmt.Errorf("error publishing: %v", err)
	}
	return nil
}
