package builtinfoo

import (
	"testing"
)

func Test_tmpValueAssignInForRange(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpValueAssignInForRange()
		})
	}
}

func Test_localValueReassignInFor(t *testing.T) {
	type args struct {
		catch bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{catch: false},
		},
		{
			"test case 2",
			args{catch: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localValueReassignInFor(tt.args.catch)
		})
	}
}

func Test_forCompareCall(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			forCompareCall(tt.args.c)
		})
	}
}
