package publisher

import (
	"cloud.google.com/go/pubsub"
	"falabella.com/boss-message-library/domain/event"
	"testing"
)

func TestNotifyOrderStatusEvent(t *testing.T) {

	type args struct {
		event *event.HistoricalEvent
	}
	tests := []struct {
		name  string
		args  args
		topic *pubsub.Topic
		want  int
	}{
		{
			"#1 notify order status event success",
			args{new(event.HistoricalEvent)},
			new(pubsub.Topic),
			0,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			rs := &EventPublisher{
				topic: tt.topic,
			}
			publish = func(t *pubsub.Topic, s interface{}, attributes map[string]string) error {
				return nil
			}

			var headers = make(map[string]string)
			headers["country"] = ""
			headers["sourceSystem"] = ""
			headers["tenantId"] = ""
			headers["version"] = "1.0"

			err := rs.Publish(tt.args.event, headers)
			if nil != err {
				t.Errorf("error when try to notify AuditOrder")
				return
			}
		})
	}
}

func TestNewOrderStatusPublisher(t *testing.T) {

	tests := []struct {
		name    string
		notWant *EventPublisher
	}{
		{
			"#1 new order status publisher success",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEventPublisher(); got == tt.notWant {
				t.Errorf("NewOrderStatusPublisher() = %v, want %v", got, tt.notWant)
			}
		})
	}
}
