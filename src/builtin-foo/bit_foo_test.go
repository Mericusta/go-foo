package builtinfoo

import (
	"testing"
)

func TestNumberEvenOddCheck(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{n: 1},
			true,
		},
		{
			"test case 2",
			args{n: 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberEvenOddCheck(tt.args.n); got != tt.want {
				t.Errorf("NumberEvenOddCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZoomInAndOutInMultiplesOf2(t *testing.T) {
	type args struct {
		n  int
		in bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{n: 8, in: true},
			4,
		},
		{
			"test case 2",
			args{n: 8, in: false},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZoomInAndOutInMultiplesOf2(tt.args.n, tt.args.in); got != tt.want {
				t.Errorf("ZoomInAndOutInMultiplesOf2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZoomOutInMultiplesOf10(t *testing.T) {
	type args struct {
		n      int
		origin bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{n: 1, origin: true},
			10,
		},
		{
			"test case 2",
			args{n: 2, origin: true},
			20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZoomOutInMultiplesOf10(tt.args.n, tt.args.origin); got != tt.want {
				t.Errorf("ZoomOutInMultiplesOf10() = %v, want %v", got, tt.want)
			}
		})
	}
}
