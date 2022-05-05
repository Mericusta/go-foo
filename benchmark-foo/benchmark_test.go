package benchmarkfoo

import (
	"math"
	"testing"
)

func Benchmark_mapFunction(b *testing.B) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"map",
			args{
				count: math.MaxInt16,
			},
		},
	}
	b.ResetTimer()
	for _, tt := range tests {
		mapFunction(tt.args.count)
	}
}

func Benchmark_sliceFunction(b *testing.B) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"slice",
			args{
				count: math.MaxInt16,
			},
		},
	}
	b.ResetTimer()
	for _, tt := range tests {
		sliceFunction(tt.args.count)
	}
}

func Benchmark_arrayFunction(b *testing.B) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"array",
			args{
				count: math.MaxInt16,
			},
		},
	}
	b.ResetTimer()
	for _, tt := range tests {
		arrayFunction(tt.args.count)
	}
}
