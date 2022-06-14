package log

import (
	"github.com/rendis/abslog"
	"reflect"
	"testing"
)

func TestGetLogger(t *testing.T) {
	tests := []struct {
		name string
		want *abslog.AbsLog
	}{
		{
			name: "TestGetLogger",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLogger(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
