package iofoo

import (
	"bytes"
	"testing"
)

func BenchmarkWriteFileFoo(b *testing.B) {
	type args struct {
		writerIndex int
	}
	tests := []struct {
		name           string
		args           args
		wantOutputFile string
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			WriteFileFoo(tt.args.writerIndex, outputFile)
		}
	}
}
