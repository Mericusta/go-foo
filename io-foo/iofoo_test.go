package iofoo

import (
	"bytes"
	"testing"
)

func TestWriteFileFoo(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputFile := &bytes.Buffer{}
			WriteFileFoo(tt.args.writerIndex, outputFile)
			if gotOutputFile := outputFile.String(); gotOutputFile != tt.wantOutputFile {
				t.Errorf("WriteFileFoo() = %v, want %v", gotOutputFile, tt.wantOutputFile)
			}
		})
	}
}
