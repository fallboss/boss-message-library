package properties

import "sync"

var (
	publisherPropsInstanceOnce sync.Once
	publisherPropsInstance *PublisherProp
)

func GetPublisherProperty() *PublisherProp {
	publisherPropsInstanceOnce.Do(func() {
		publisherPropsInstance = &PublisherProp{}
	})

	return publisherPropsInstance
}

// PublisherProp Property Struct Definition
type PublisherProp struct {
	Publisher `yaml:"publishers"`
}

type Publisher struct {
	OrderStatus PubInfo `yaml:"historical-event"`
}

type PubInfo struct {
	ProjectId string `yaml:"project-id"`
	TopicId   string `yaml:"topic-id"`
}
