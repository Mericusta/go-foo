package arrayfoo

import (
	"testing"
)

func TestClearArrayFoo(t *testing.T) {
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
			ClearArrayFoo()
		})
	}
}

func TestReturnArrayBeforeIndexFoo(t *testing.T) {
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
			ReturnArrayBeforeIndexFoo()
		})
	}
}
