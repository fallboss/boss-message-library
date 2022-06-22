package properties

import (
	"falabella.com/boss-message-library/publisher"
	"reflect"
	"testing"
)

func TestGetPublisherProperty(t *testing.T) {
	tests := []struct {
		name string
		want *publisher.PublisherProp
	}{
		{
			name: "test",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := publisher.GetPublisherProperty(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPublisherProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}
