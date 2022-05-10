package benchmarkfoo

import (
	"fmt"
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

// ----------------------------------------------------------------

func Benchmark_Pray(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"benchmark",
		},
	}
	b.ResetTimer()
	for index := 0; index != len(tests); index++ {
		m := Pray()
		fmt.Printf("%v\n", m)
	}
	b.StopTimer()
}

func Benchmark_PrayOpt(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"benchmark_opt",
		},
	}
	b.ResetTimer()
	for index := 0; index != len(tests); index++ {
		s := PrayOpt()
		fmt.Printf("%v\n", s)
	}
	b.StopTimer()
}

func TestPray(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Pray()
		})
	}
}

func TestLambdaCapture(t *testing.T) {
	type args struct {
		testCase int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "lambda capture",
			args: args{
				testCase: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LambdaCapture(tt.args.testCase); got != tt.want {
				t.Errorf("LambdaCapture() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLambdaCapture(b *testing.B) {
	type args struct {
		testCase int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "lambda capture",
			args: args{
				testCase: 1,
			},
			want: 0,
		},
	}
	b.ResetTimer()
	for _, tt := range tests {
		v := LambdaCapture(tt.args.testCase)
		fmt.Printf("v = %v\n", v)
	}
}
