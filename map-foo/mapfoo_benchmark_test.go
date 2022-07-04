package mapfoo

import (
	"testing"
)

func BenchmarkMapCapacityFoo(b *testing.B) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1: c 4",
			args{c: 16},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			MapCapacityFoo(tt.args.c)
		}
	}
}
