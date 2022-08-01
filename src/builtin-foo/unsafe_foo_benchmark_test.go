package builtinfoo

import (
	"math"
	"testing"
)

func BenchmarkBytesToString(b *testing.B) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"test case abcdefu",
			args{b: []byte{97, 98, 99, 100, 101, 102, 117}},
			"abcdefu",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			BytesToString(tt.args.b)
		}
	}
}

func BenchmarkBytesToStringFool(b *testing.B) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"test case abcdefu",
			args{b: []byte{97, 98, 99, 100, 101, 102, 117}},
			"abcdefu",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			BytesToStringFool(tt.args.b)
		}
	}
}

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

func BenchmarkStringToBytes(b *testing.B) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			"test case abcdefu",
			args{s: "abcdefu"},
			[]byte{97, 98, 99, 100, 101, 102, 117},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			StringToBytes(tt.args.s)
		}
	}
}

func BenchmarkStringToBytesFool(b *testing.B) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			"test case abcdefu",
			args{s: "abcdefu"},
			[]byte{97, 98, 99, 100, 101, 102, 117},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			StringToBytesFool(tt.args.s)
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
