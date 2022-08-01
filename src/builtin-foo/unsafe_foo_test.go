package builtinfoo

import (
	"math"
	"reflect"
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

func TestBytesToString(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToString(tt.args.b); got != tt.want {
				t.Errorf("BytesToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToBytes(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToBytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToStringFool(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToStringFool(tt.args.b); got != tt.want {
				t.Errorf("BytesToStringFool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToBytesFool(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToBytesFool(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToBytesFool() = %v, want %v", got, tt.want)
			}
		})
	}
}
