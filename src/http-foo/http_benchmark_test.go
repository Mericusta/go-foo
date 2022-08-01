package httpfoo

import (
	"testing"
)

func BenchmarkRequestExample(b *testing.B) {
	type args struct {
		index int
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
			RequestExample(tt.args.index)
		}
	}
}
