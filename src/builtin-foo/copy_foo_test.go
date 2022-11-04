package builtinfoo

import (
	"reflect"
	"testing"
)

func TestCopyStringSliceFromStringsSplit(t *testing.T) {
	type args struct {
		c       int
		s       string
		lenCase string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			"test case 1, equal",
			args{c: 3, s: "a1,b2,c3", lenCase: "=="},
			[]string{"a1", "b2", "c3"},
		},
		{
			"test case 1, equal",
			args{c: 3, s: "a1,b2,c3", lenCase: "<"},
			[]string{"a1", "b2"},
		},
		{
			"test case 1, equal",
			args{c: 3, s: "a1,b2,c3", lenCase: ">"},
			[]string{"a1", "b2", "c3", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CopyStringSliceFromStringsSplit(tt.args.c, tt.args.s, tt.args.lenCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopyStringSliceFromStringsSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyByteSliceFromStringWithThreeCases(t *testing.T) {
	type args struct {
		c       int
		s       string
		lenCase string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			"test case 1, equal",
			args{c: 8, s: "a1,b2,c3", lenCase: "=="},
			[]byte("a1,b2,c3"),
		},
		{
			"test case 1, equal",
			args{c: 8, s: "a1,b2,c3", lenCase: "<"},
			[]byte("a1,b2,c"),
		},
		{
			"test case 1, equal",
			args{c: 8, s: "a1,b2,c3", lenCase: ">"},
			append([]byte("a1,b2,c3"), 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CopyByteSliceFromStringWithThreeCases(tt.args.c, tt.args.s, tt.args.lenCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopyByteSliceFromStringWithThreeCases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyByteSliceFromString(t *testing.T) {
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
			"test case 1",
			args{s: "a1,b2,c3,d4,e5,f6,g8"},
			[]byte("a1,b2,c3,d4,e5,f6,g8"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CopyByteSliceFromString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopyByteSliceFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyMyself(t *testing.T) {
	type args struct {
		b    []int
		from int
		to   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				b: func() []int {
					b := make([]int, 0, 16)
					for index := 0; index != 16; index++ {
						b = append(b, index)
					}
					return b
				}(),
				from: 8,
				to:   16,
			},
			[]int{
				8, 9, 10, 11, 12, 13, 14, 15,
				8, 9, 10, 11, 12, 13, 14, 15,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CopyMyself(tt.args.b, tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopyMyself() = %v, want %v", got, tt.want)
			}
		})
	}
}
