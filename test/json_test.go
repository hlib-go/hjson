package test

import (
	"testing"
)

func TestJsonToMap(t *testing.T) {
	type args struct {
		data string
		v    interface{}
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

func TestJsonToStruct(t *testing.T) {
	type args struct {
		data string
		v    interface{}
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
