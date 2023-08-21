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
		url              string
		header           map[string]string
		d                time.Duration
		useResty         bool
		concurrency      bool
		concurrencyCount int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				url:      "http://127.0.0.1:8182/pay/cb/mock",
				header:   map[string]string{"Origin": "http://ios.appstore.com"},
				d:        0,
				useResty: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JustPost(tt.args.url, tt.args.header, tt.args.d, tt.args.useResty, tt.args.concurrency, tt.args.concurrencyCount)
		})
	}
}
