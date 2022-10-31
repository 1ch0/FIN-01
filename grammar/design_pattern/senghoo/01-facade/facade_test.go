package _1_facade

import (
	"reflect"
	"testing"
)

func TestNewAPI(t *testing.T) {
	tests := []struct {
		name string
		want Test
	}{
		{
			name: "test",
			want: &apiImpl{
				a: NewTestA(),
				b: NewTestB(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAPI(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAPI() = %v, want %v", got, tt.want)
			} else {
				t.Log(got.Test())
			}
		})
	}
}
