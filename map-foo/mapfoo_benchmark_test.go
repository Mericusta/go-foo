package mapfoo

import (
	"testing"
)

func BenchmarkMapCapacityFoo(b *testing.B) {
	type args struct {
		count    int
		capacity int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case",
			args{count: 23, capacity: 38},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			MapCapacityFoo(tt.args.count, tt.args.capacity)
		}
	}
}
