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

func Test_BitOp1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"test case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BitOp1()
		})
	}
}

func Test_BitOp2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"test case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BitOp2()
		})
	}
}
