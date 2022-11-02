package builtinfoo

import (
	"testing"
)

func BenchmarkNumberEvenOddCheck(b *testing.B) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{n: 1},
			true,
		},
		{
			"test case 2",
			args{n: 2},
			false,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			NumberEvenOddCheck(tt.args.n)
		}
	}
}

func BenchmarkZoomInAndOutInMultiplesOf2(b *testing.B) {
	type args struct {
		n  int
		in bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{n: 8, in: true},
			4,
		},
		{
			"test case 2",
			args{n: 8, in: false},
			16,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			ZoomInAndOutInMultiplesOf2(tt.args.n, tt.args.in)
		}
	}
}

func BenchmarkZoomOutInMultiplesOf10(b *testing.B) {
	type args struct {
		n      int
		origin bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{n: 1, origin: true},
			10,
		},
		{
			"test case 2",
			args{n: 2, origin: true},
			20,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			ZoomOutInMultiplesOf10(tt.args.n, tt.args.origin)
		}
	}
}
