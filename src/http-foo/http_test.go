package httpfoo

import (
	"testing"
	"time"
)

func TestRequestExample(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{index: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RequestExample(tt.args.index)
		})
	}
}

func TestJustPost(t *testing.T) {
	type args struct {
		d        time.Duration
		useResty bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				d:        0,
				useResty: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JustPost(tt.args.d, tt.args.useResty)
		})
	}
}
