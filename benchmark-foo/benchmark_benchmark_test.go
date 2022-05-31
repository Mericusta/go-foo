package benchmarkfoo

import (
	"testing"
)

func BenchmarkFMTPrintfBenchmark(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			FMTPrintfBenchmark()
		}
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
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			LambdaCapture(tt.args.testCase)
		}
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
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			ElementsInSlice(tt.args.count)
		}
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
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			ElementsInMap(tt.args.count)
		}
	}
}
