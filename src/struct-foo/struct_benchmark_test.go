package structfoo

import (
	"testing"
)

func BenchmarkBaseStructTrace(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			BaseStructTrace()
		}
	}
}

func BenchmarkDerivativeWithPointerBase(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			DerivativeWithPointerBase()
		}
	}
}

func BenchmarkStructThisMemberDiff(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			StructThisMemberDiff()
		}
	}
}

func BenchmarkSubStructAssign(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			SubStructAssign()
		}
	}
}

func BenchmarkSubStructDerivative(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			SubStructDerivative()
		}
	}
}

func BenchmarkSwapStructValueOneLine(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			SwapStructValueOneLine()
		}
	}
}

func Benchmark_base_Input(b *testing.B) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		b    *base
		args args
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			tt.b.Input(tt.args.i)
		}
	}
}

func Benchmark_base_Output(b *testing.B) {
	tests := []struct {
		name string
		b    *base
		want int
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			tt.b.Output()
		}
	}
}

func Benchmark_derivative_Input(b *testing.B) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		b    *derivative
		args args
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			tt.b.Input(tt.args.i)
		}
	}
}

func Benchmark_derivative_ModBMap(b *testing.B) {
	type args struct {
		k int
		v int
	}
	tests := []struct {
		name string
		b    *derivative
		args args
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			tt.b.ModBMap(tt.args.k, tt.args.v)
		}
	}
}

func Benchmark_derivative_Output(b *testing.B) {
	tests := []struct {
		name string
		d    *derivative
		want int
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			tt.d.Output()
		}
	}
}

func Benchmark_newBase(b *testing.B) {
	tests := []struct {
		name string
		want base
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			newBase()
		}
	}
}

func Benchmark_newBasePointer(b *testing.B) {
	tests := []struct {
		name string
		want *base
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			newBasePointer()
		}
	}
}

func Benchmark_stmd_GetCopyThisV(b *testing.B) {
	tests := []struct {
		name  string
		s     stmd
		want  map[int]int
		want1 []int
		want2 int
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			tt.s.GetCopyThisV()
		}
	}
}

func Benchmark_stmd_GetPointerThisV(b *testing.B) {
	tests := []struct {
		name  string
		s     *stmd
		want  map[int]int
		want1 []int
		want2 int
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			tt.s.GetPointerThisV()
		}
	}
}
