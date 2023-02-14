package algorithmfoo

import (
	"testing"
)

func BenchmarkPutAwayRandom(b *testing.B) {
	type args struct {
		count     int
		rangeFunc func(func(int, int) bool)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			PutAwayRandom(tt.args.count, tt.args.rangeFunc)
		}
	}
}

func BenchmarkPutBackRandom(b *testing.B) {
	type args struct {
		count     int
		rangeFunc func(func(int, int) bool)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			PutBackRandom(tt.args.count, tt.args.rangeFunc)
		}
	}
}
