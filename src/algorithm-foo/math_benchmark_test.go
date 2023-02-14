package algorithmfoo

import (
	"math"
	"testing"
)

func BenchmarkCalculateNumberLenByCycling(b *testing.B) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{n: math.MaxInt64},
			19,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			CalculateNumberLenByCycling(tt.args.n)
		}
	}
}

func BenchmarkCalculateNumberLenByLog(b *testing.B) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{n: math.MaxInt64},
			19,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			CalculateNumberLenByLog(tt.args.n)
		}
	}
}

func BenchmarkCalculateNumberLenByRecursive(b *testing.B) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{n: math.MaxInt64},
			19,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			CalculateNumberLenByRecursive(tt.args.n)
		}
	}
}
