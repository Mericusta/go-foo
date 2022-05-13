package benchmarkfoo

import (
	"fmt"
	"math"
	"reflect"
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

func TestElementsInSlice(t *testing.T) {
	getEvenSlice := func(c int) []int {
		s := make([]int, 0, c/2)
		for index := 0; index < c; index++ {
			if index%2 == 0 {
				s = append(s, index)
			}
		}
		return s
	}
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"8 elements slice test",
			args{
				count: 8,
			},
			getEvenSlice(8),
		},
		{
			"16 elements slice test",
			args{
				count: 16,
			},
			getEvenSlice(16),
		},
		{
			"32 elements slice test",
			args{
				count: 32,
			},
			getEvenSlice(32),
		},
		{
			"64 elements slice test",
			args{
				count: 64,
			},
			getEvenSlice(64),
		},
		{
			"128 elements slice test",
			args{
				count: 128,
			},
			getEvenSlice(128),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ElementsInSlice(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementsInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkElementsInSlice(b *testing.B) {
	getEvenSlice := func(c int) []int {
		s := make([]int, 0, c/2)
		for index := 0; index < c; index++ {
			if index%2 == 0 {
				s = append(s, index)
			}
		}
		return s
	}
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"8 elements slice test",
			args{
				count: 8,
			},
			getEvenSlice(8),
		},
		{
			"16 elements slice test",
			args{
				count: 16,
			},
			getEvenSlice(16),
		},
		{
			"32 elements slice test",
			args{
				count: 32,
			},
			getEvenSlice(32),
		},
		{
			"64 elements slice test",
			args{
				count: 64,
			},
			getEvenSlice(64),
		},
		{
			"128 elements slice test",
			args{
				count: 128,
			},
			getEvenSlice(128),
		},
	}
	b.ResetTimer()
	for _, tt := range tests {
		ElementsInSlice(tt.args.count)
	}
}

func TestElementsInMap(t *testing.T) {
	getEvenMap := func(c int) map[int]struct{} {
		m := make(map[int]struct{}, c/2)
		for index := 0; index < c; index++ {
			if index%2 == 0 {
				m[index] = struct{}{}
			}
		}
		return m
	}
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want map[int]struct{}
	}{
		// TODO: Add test cases.
		{
			"8 elements map test",
			args{
				count: 8,
			},
			getEvenMap(8),
		},
		{
			"16 elements map test",
			args{
				count: 16,
			},
			getEvenMap(16),
		},
		{
			"32 elements map test",
			args{
				count: 32,
			},
			getEvenMap(32),
		},
		{
			"64 elements map test",
			args{
				count: 64,
			},
			getEvenMap(64),
		},
		{
			"128 elements map test",
			args{
				count: 128,
			},
			getEvenMap(128),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ElementsInMap(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementsInMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkElementsInMap(b *testing.B) {
	getEvenMap := func(c int) map[int]struct{} {
		m := make(map[int]struct{}, c/2)
		for index := 0; index < c; index++ {
			if index%2 == 0 {
				m[index] = struct{}{}
			}
		}
		return m
	}
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want map[int]struct{}
	}{
		// TODO: Add test cases.
		{
			"8 elements map test",
			args{
				count: 8,
			},
			getEvenMap(8),
		},
		{
			"16 elements map test",
			args{
				count: 16,
			},
			getEvenMap(16),
		},
		{
			"32 elements map test",
			args{
				count: 32,
			},
			getEvenMap(32),
		},
		{
			"64 elements map test",
			args{
				count: 64,
			},
			getEvenMap(64),
		},
		{
			"128 elements map test",
			args{
				count: 128,
			},
			getEvenMap(128),
		},
	}
	b.ResetTimer()
	for _, tt := range tests {
		ElementsInMap(tt.args.count)
	}
}
