package slicefoo

import (
	"testing"
)

func BenchmarkPassSliceAndChangeIt(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			PassSliceAndChangeIt()
		}
	}
}

func Benchmark_appendGreaterCapacitySliceFunc(b *testing.B) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			appendGreaterCapacitySliceFunc(tt.args.s)
		}
	}
}

func Benchmark_appendLowerCapacitySliceFunc(b *testing.B) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			appendLowerCapacitySliceFunc(tt.args.s)
		}
	}
}

func Benchmark_updateSliceFunc(b *testing.B) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			updateSliceFunc(tt.args.s)
		}
	}
}
