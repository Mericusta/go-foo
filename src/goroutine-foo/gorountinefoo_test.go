package goroutinefoo

import (
	"testing"
)

func TestOpenSoMuchGoRoutine(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OpenSoMuchGoRoutine()
		})
	}
}

func TestAllMIsWorking(t *testing.T) {
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
			AllMIsWorking()
		})
	}
}
