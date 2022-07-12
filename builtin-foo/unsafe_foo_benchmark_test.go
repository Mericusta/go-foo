package builtinfoo

import (
	"math"
	"testing"
)

func BenchmarkModifyPrivateValue(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			ModifyPrivateValue()
		}
	}
}

func BenchmarkTraversalSliceByFor(b *testing.B) {
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			TraversalSliceByFor(tt.args.data)
		}
	}
}

func BenchmarkTraversalSliceByForRange(b *testing.B) {
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			TraversalSliceByForRange(tt.args.data)
		}
	}
}

func BenchmarkTraversalSliceByUsingUnsafePointer(b *testing.B) {
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			TraversalSliceByUsingUnsafePointer(tt.args.data)
		}
	}
}
