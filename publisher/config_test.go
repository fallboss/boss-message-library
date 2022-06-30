package publisher

import (
	"reflect"
	"testing"
)

func TestGetPublisherProperty(t *testing.T) {
	tests := []struct {
		name string
		want *PublisherProp
	}{
		{
			name: "test",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPublisherProperty(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPublisherProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}
