package main

import "testing"

func Test_dialFoo(t *testing.T) {
	type args struct {
		urlStr string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{urlStr: "ws://192.168.2.203:6666/dial"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dialFoo(tt.args.urlStr)
		})
	}
}
