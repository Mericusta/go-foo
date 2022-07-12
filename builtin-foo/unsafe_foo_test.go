package builtinfoo

import (
	"math"
	"testing"
)

func TestTraversalSliceByUsingUnsafePointer(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case: math.MaxInt16",
			args{[]byte(func() []byte {
				s := make([]byte, math.MaxInt16)
				for i := 0; i != math.MaxInt16; i++ {
					s[i] = byte('a' + i%26)
				}
				return s
			}())},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TraversalSliceByUsingUnsafePointer(tt.args.data)
		})
	}
}

func TestTraversalSliceByForRange(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case: math.MaxInt16",
			args{[]byte(func() []byte {
				s := make([]byte, math.MaxInt16)
				for i := 0; i != math.MaxInt16; i++ {
					s[i] = byte('a' + i%26)
				}
				return s
			}())},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TraversalSliceByForRange(tt.args.data)
		})
	}
}

func TestTraversalSliceByFor(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case: math.MaxInt16",
			args{[]byte(func() []byte {
				s := make([]byte, math.MaxInt16)
				for i := 0; i != math.MaxInt16; i++ {
					s[i] = byte('a' + i%26)
				}
				return s
			}())},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TraversalSliceByFor(tt.args.data)
		})
	}
}

func TestModifyPrivateValue(t *testing.T) {
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
			ModifyPrivateValue()
		})
	}
}
