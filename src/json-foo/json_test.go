package jsonfoo

import (
	"testing"
)

func TestJsonFoo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JsonFoo()
		})
	}
}

func TestInterfaceMarshalFoo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InterfaceMarshalFoo()
		})
	}
}

func TestSliceMarshalFoo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SliceMarshalFoo()
		})
	}
}
