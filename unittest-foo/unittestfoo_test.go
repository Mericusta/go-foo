package unittestfoo

import (
	"reflect"
	"testing"
)

func TestNewPointerFunc(t *testing.T) {
	tests := []struct {
		name string
		want *AnExampleStruct
	}{
		// TODO: Add test cases.
		{
			"example 1",
			&AnExampleStruct{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPointerFunc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPointerFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPointerMapFunc(t *testing.T) {
	tests := []struct {
		name string
		want map[int]*AnExampleStruct
	}{
		// TODO: Add test cases.
		{
			"example 1",
			map[int]*AnExampleStruct{
				1: {v: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPointerMapFunc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPointerMapFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
