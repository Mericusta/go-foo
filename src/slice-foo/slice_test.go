package slicefoo

import (
	"testing"
)

func Test_appendLowerCapacitySliceFunc(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appendLowerCapacitySliceFunc(tt.args.s)
		})
	}
}

func Test_updateSliceFunc(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateSliceFunc(tt.args.s)
		})
	}
}

func Test_appendGreaterCapacitySliceFunc(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appendGreaterCapacitySliceFunc(tt.args.s)
		})
	}
}

func TestPassSliceAndChangeIt(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PassSliceAndChangeIt()
		})
	}
}

func TestResetSliceFoo(t *testing.T) {
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
			ResetSliceFoo()
		})
	}
}
