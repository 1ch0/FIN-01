package _0_simple_factory

import (
	"reflect"
	"testing"
)

func TestAPI(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want API
	}{
		{
			name: "hi",
			args: args{
				name: "hi",
			},
			want: &hiAPI{},
		},
		{
			name: "hello",
			args: args{
				name: "hello",
			},
			want: &helloAPI{},
		},
		{
			name: "hi",
			args: args{
				name: "",
			},
			want: &hiAPI{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAPI(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAPI() = %v, want %v", got, tt.want)
			} else {
				t.Log(got.Say("test"))
			}
		})
	}
}
