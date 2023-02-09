package benchmarkfoo

import (
	"testing"
)

func Benchmark_removeDuplication_map(b *testing.B) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				arr: []string{
					"test_case_1",
					"test", "case", "1",
					"case_1", "test",
					"test_1", "case",
					"test_case", "1",
				},
			},
			[]string{
				"test_case_1",
				"test", "case", "1",
				"case_1", "test_1", "test_case",
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			removeDuplication_map(tt.args.arr)
		}
	}
}

func Benchmark_removeDuplication_sort(b *testing.B) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				arr: []string{
					"test_case_1",
					"test", "case", "1",
					"case_1", "test",
					"test_1", "case",
					"test_case", "1",
				},
			},
			[]string{
				"test_case_1",
				"test", "case", "1",
				"case_1", "test_1", "test_case",
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			removeDuplication_sort(tt.args.arr)
		}
	}
}

func Benchmark_simple_removeDuplication_map(b *testing.B) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				arr: []string{
					"test_case_1",
					"test", "case", "1",
					"case_1", "test",
					"test_1", "case",
					"test_case", "1",
				},
			},
			[]string{
				"test_case_1",
				"test", "case", "1",
				"case_1", "test_1", "test_case",
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			simple_removeDuplication_map(tt.args.arr)
		}
	}
}
