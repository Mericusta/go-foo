package algorithmfoo

import (
	"testing"
)

func BenchmarkDecodeID(b *testing.B) {
	type args struct {
		identifier int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			"test case 740307149",
			args{identifier: 740307149},
			1024,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			DecodeID(tt.args.identifier)
		}
	}
}

func BenchmarkDynamicHashAverageAlgorithm(b *testing.B) {
	type args struct {
		nCount int
		sCount int
		rCount int
		luckyS int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				nCount: 10000,
				sCount: 7,
				rCount: 1,
				luckyS: 5,
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			DynamicHashAverageAlgorithm(tt.args.nCount, tt.args.sCount, tt.args.rCount, tt.args.luckyS)
		}
	}
}

func BenchmarkEncodeID(b *testing.B) {
	type args struct {
		ID int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			"test case 1024",
			args{ID: 1024},
			740307149,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			EncodeID(tt.args.ID)
		}
	}
}

func Benchmark_modHash(b *testing.B) {
	type args struct {
		n int
		s int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			modHash(tt.args.n, tt.args.s)
		}
	}
}
