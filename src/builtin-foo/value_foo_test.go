package builtinfoo

import "testing"

func TestFloat32ZeroDivide(t *testing.T) {
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
			Float32ZeroDivide()
		})
	}
}
