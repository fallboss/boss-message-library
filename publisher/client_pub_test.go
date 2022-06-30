package publisher

import (
	"cloud.google.com/go/pubsub"
	"reflect"
	"testing"
)

func Test_getClient(t *testing.T) {
	type args struct {
		projectId string
	}
	tests := []struct {
		name string
		args args
		want *pubsub.Client
	}{
		{
			name: "getClient",
			args: args{
				projectId: "test",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getClient(tt.args.projectId); reflect.DeepEqual(got, tt.want) {
				t.Errorf("getClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildTopic(t *testing.T) {
	type args struct {
		projectId string
		topicId   string
	}
	tests := []struct {
		name string
		args args
		want *pubsub.Topic
	}{
		{
			name: "BuildTopic",
			args: args{
				projectId: "test",
				topicId:   "test",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildTopic(tt.args.projectId, tt.args.topicId); reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTopic() = %v, want %v", got, tt.want)
			}
		})
	}
}
