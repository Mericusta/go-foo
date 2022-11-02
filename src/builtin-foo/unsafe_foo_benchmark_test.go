package builtinfoo

import (
	"math"
	"strings"
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
		// {
		// 	// BenchmarkStringToBytes-12    	233983488	         4.993 ns/op	       0 B/op	       0 allocs/op
		// 	"test case a1,b2,c3,d4,e5,f6,g8",
		// 	args{s: "a1,b2,c3,d4,e5,f6,g8"},
		// 	[]byte("a1,b2,c3,d4,e5,f6,g8"),
		// },
		{
			// BenchmarkStringToBytes-12    	708182278	         1.618 ns/op	       0 B/op	       0 allocs/op
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
		// {
		// 	// BenchmarkStringToBytesFool-12    	250087009	         5.033 ns/op	       0 B/op	       0 allocs/op
		// 	"test case a1,b2,c3,d4,e5,f6,g8",
		// 	args{s: "a1,b2,c3,d4,e5,f6,g8"},
		// 	[]byte("a1,b2,c3,d4,e5,f6,g8"),
		// },
		{
			// BenchmarkStringToBytesFool-12    	 5085036	       250.2 ns/op	    1280 B/op	       1 allocs/op
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
