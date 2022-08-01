package builtinfoo

import (
	"reflect"
	"testing"
)

func TestSwitchCaseValueFallthrough(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{c: 5},
			[]int{5, 6, 7, 8, 9},
		},
		{
			"test case 9",
			args{c: 9},
			[]int{9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SwitchCaseValueFallthrough(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwitchCaseValueFallthrough() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwitchCaseExpressionFallthrough(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{c: 51},
			[]int{5, 6, 7, 8, 9},
		},
		{
			"test case 9",
			args{c: 91},
			[]int{9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SwitchCaseExpressionFallthrough(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwitchCaseExpressionFallthrough() = %v, want %v", got, tt.want)
			}
		})
	}
}
