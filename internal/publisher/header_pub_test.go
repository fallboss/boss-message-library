package publisher

import "testing"

func TestCompleteHeader(t *testing.T) {
	type args struct {
		attributes map[string]string
		eventType  string
		entityType string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Test",
			args{
				make(map[string]string),
				"eventType",
				"entityType",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CompleteHeader(tt.args.attributes, tt.args.eventType, tt.args.entityType)
			if tt.args.attributes["eventType"] != "eventType" {
				t.Errorf("CompleteHeader() = %v, want %v", tt.args.attributes["eventType"], "eventType")
			}
			if tt.args.attributes["entityType"] != "entityType" {
				t.Errorf("CompleteHeader() = %v, want %v", tt.args.attributes["entityType"], "entityType")
			}
		})
	}
}
