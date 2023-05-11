package structfoo

import (
	"reflect"
	"testing"
)

func TestSwapStructValueOneLine(t *testing.T) {
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
			SwapStructValueOneLine()
		})
	}
}

func Test_stmd_GetPointerThisV(t *testing.T) {
	tests := []struct {
		name  string
		s     *stmd
		want  map[int]int
		want1 []int
		want2 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.GetPointerThisV()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stmd.GetPointerThisV() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("stmd.GetPointerThisV() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("stmd.GetPointerThisV() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_stmd_GetCopyThisV(t *testing.T) {
	tests := []struct {
		name  string
		s     stmd
		want  map[int]int
		want1 []int
		want2 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.s.GetCopyThisV()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stmd.GetCopyThisV() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("stmd.GetCopyThisV() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("stmd.GetCopyThisV() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestStructThisMemberDiff(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StructThisMemberDiff()
		})
	}
}

func Test_base_Output(t *testing.T) {
	tests := []struct {
		name string
		b    *base
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Output(); got != tt.want {
				t.Errorf("base.Output() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base_Input(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		b    *base
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.Input(tt.args.i)
		})
	}
}

func Test_derivative_Output(t *testing.T) {
	tests := []struct {
		name string
		d    *derivative
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Output(); got != tt.want {
				t.Errorf("derivative.Output() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_derivative_Input(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		b    *derivative
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.Input(tt.args.i)
		})
	}
}

func Test_derivative_ModBMap(t *testing.T) {
	type args struct {
		k int
		v int
	}
	tests := []struct {
		name string
		b    *derivative
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.ModBMap(tt.args.k, tt.args.v)
		})
	}
}

func TestDerivativeWithPointerBase(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DerivativeWithPointerBase()
		})
	}
}

func Test_newBase(t *testing.T) {
	tests := []struct {
		name string
		want base
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newBase(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newBasePointer(t *testing.T) {
	tests := []struct {
		name string
		want *base
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newBasePointer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newBasePointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseStructTrace(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BaseStructTrace()
		})
	}
}

func TestSubStructAssign(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SubStructAssign()
		})
	}
}

func TestSubStructDerivative(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SubStructDerivative()
		})
	}
}
