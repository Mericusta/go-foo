package algorithmfoo

import (
	"testing"
)

func TestEncodeID(t *testing.T) {
	type args struct {
		ID int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			"test case 1024",
			args{ID: 1024},
			740307149,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeID(tt.args.ID); got != tt.want {
				t.Errorf("EncodeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeID(t *testing.T) {
	type args struct {
		identifier int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			"test case 740307149",
			args{identifier: 740307149},
			1024,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeID(tt.args.identifier); got != tt.want {
				t.Errorf("DecodeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamicHashAverageAlgorithm(t *testing.T) {
	type args struct {
		nCount int
		sCount int
		rCount int
		luckyS int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				nCount: 10000,
				sCount: 7,
				rCount: 1,
				luckyS: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DynamicHashAverageAlgorithm(tt.args.nCount, tt.args.sCount, tt.args.rCount, tt.args.luckyS)
		})
	}
}

func Test_modHash(t *testing.T) {
	type args struct {
		n int
		s int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modHash(tt.args.n, tt.args.s); got != tt.want {
				t.Errorf("modHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
