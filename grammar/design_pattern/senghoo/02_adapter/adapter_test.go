package _2_adapter

import (
	"reflect"
	"testing"
)

func TestNewAdapter(t *testing.T) {
	tests := []struct {
		name string
		want Target
	}{
		{
			"adapter",
			adapterImpl{adapteeImpl{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAdapter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAdapter() = %v, want %v", got, tt.want)
			} else {
				t.Log(got.Request())
			}
		})
	}
}
