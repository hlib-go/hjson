package test

import (
	"testing"

	"github.com/mszsgo/hjson"
)

func TestMapToJson(t *testing.T) {
	type args struct {
		mapData interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hjson.MapToJson(tt.args.mapData); got != tt.want {
				t.Errorf("MapToJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapToStruct(t *testing.T) {
	type args struct {
		mapData interface{}
		v       interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
