package builtinfoo

import (
	"testing"
)

func BenchmarkSwitchCaseExpressionFallthrough(b *testing.B) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{c: 51},
			[]int{5, 6, 7, 8, 9},
		},
		{
			"test case 9",
			args{c: 91},
			[]int{9},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			SwitchCaseExpressionFallthrough(tt.args.c)
		}
	}
}

func BenchmarkSwitchCaseValueFallthrough(b *testing.B) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{c: 5},
			[]int{5, 6, 7, 8, 9},
		},
		{
			"test case 9",
			args{c: 9},
			[]int{9},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			SwitchCaseValueFallthrough(tt.args.c)
		}
	}
}
