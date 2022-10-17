package iofoo

import (
	"bytes"
	"math"
	"reflect"
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

func Test_tlvSocketPacketFoo(t *testing.T) {
	type args struct {
		len   int
		value uint32
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "test case 1: put max uint8",
			args: args{len: 4, value: math.MaxUint8},
			want: []byte{0, 0, 0, 255},
		},
		{
			name: "test case 2: put max uint16",
			args: args{len: 4, value: math.MaxUint16},
			want: []byte{0, 0, 255, 255},
		},
		{
			name: "test case 3: put max uint24",
			args: args{len: 4, value: (math.MaxUint16 << 8) | math.MaxUint8},
			want: []byte{0, 255, 255, 255},
		},
		{
			name: "test case 4: put max uint32",
			args: args{len: 4, value: math.MaxUint32},
			want: []byte{255, 255, 255, 255},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tlvSocketPacketFoo(tt.args.len, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tlvSocketPacketFoo() = %v, want %v", got, tt.want)
			}
		})
	}
}
