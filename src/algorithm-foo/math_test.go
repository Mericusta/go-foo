package algorithmfoo

import (
	"math"
	"testing"
)

func TestCalculateNumberLenByCycling(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{n: math.MaxInt64},
			19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateNumberLenByCycling(tt.args.n); got != tt.want {
				t.Errorf("CalculateNumberLenByCycling() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateNumberLenByRecursive(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{n: math.MaxInt64},
			19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateNumberLenByRecursive(tt.args.n); got != tt.want {
				t.Errorf("CalculateNumberLenByRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateNumberLenByLog(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{n: math.MaxInt64},
			19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateNumberLenByLog(tt.args.n); got != tt.want {
				t.Errorf("CalculateNumberLenByLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateDigits(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{v: 9},
		},
		{
			"test case 1",
			args{v: 99},
		},
		{
			"test case 1",
			args{v: 999},
		},
		{
			"test case 1",
			args{v: 9999},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CalculateDigits(tt.args.v)
		})
	}
}
