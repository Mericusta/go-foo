package builtinfoo

import (
	"strings"
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
		// {
		// 	// BenchmarkCopyByteSliceFromString-12    	62855854	        20.10 ns/op	      24 B/op	       1 allocs/op
		// 	"test case 1",
		// 	args{s: "a1,b2,c3,d4,e5,f6,g8"},
		// 	[]byte("a1,b2,c3,d4,e5,f6,g8"),
		// },
		{
			// BenchmarkCopyByteSliceFromString-12    	 5114528	       248.1 ns/op	    1280 B/op	       1 allocs/op
			"test case long string length greater than 1024",
			args{s: func() string {
				b := strings.Builder{}
				for index := 0; index != 26; index++ {
					b.WriteRune(rune('a' + index))
				}
				for index := 0; index != 26; index++ {
					b.WriteRune(rune('A' + index))
				}
				for index := 0; index != 10; index++ {
					b.WriteRune(rune('0' + index))
				}
				return strings.Repeat(b.String(), 20)
			}()},
			func() []byte {
				b := strings.Builder{}
				for index := 0; index != 26; index++ {
					b.WriteRune(rune('a' + index))
				}
				for index := 0; index != 26; index++ {
					b.WriteRune(rune('A' + index))
				}
				for index := 0; index != 10; index++ {
					b.WriteRune(rune('0' + index))
				}
				return []byte(strings.Repeat(b.String(), 20))
			}(),
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
