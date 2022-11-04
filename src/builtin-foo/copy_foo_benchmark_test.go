package builtinfoo

import (
	"testing"
)

func BenchmarkCopyByteSliceFromString(b *testing.B) {
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			CopyByteSliceFromString(tt.args.s)
		}
	}
}

func BenchmarkCopyByteSliceFromStringWithThreeCases(b *testing.B) {
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			CopyByteSliceFromStringWithThreeCases(tt.args.c, tt.args.s, tt.args.lenCase)
		}
	}
}

func BenchmarkCopyMyself(b *testing.B) {
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			CopyMyself(tt.args.b, tt.args.from, tt.args.to)
		}
	}
}

func BenchmarkCopyStringSliceFromStringsSplit(b *testing.B) {
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			CopyStringSliceFromStringsSplit(tt.args.c, tt.args.s, tt.args.lenCase)
		}
	}
}
