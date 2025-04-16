package genericfoo

import (
	"reflect"
	"testing"
)

func TestInitAllWithMap(t *testing.T) {
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
			InitAllWithMap()
		})
	}
}

func TestInitAll(t *testing.T) {
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
			InitAll()
		})
	}
}

func TestGetWithMap(t *testing.T) {
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
			GetWithMap()
		})
	}
}

func TestGet(t *testing.T) {
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
			Get()
		})
	}
}

func Test_generalInterfaceCall(t *testing.T) {
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
			generalInterfaceCall()
		})
	}
}

func Test_structPointerTypeTraitFoo_01937c1a3706a7ad9057ecd022bc074e(t *testing.T) {
	type args struct {
		in *int
	}
	tests := []struct {
		name  string
		args  args
		want0 int
	}{
		{
			"test case 1",
			args{
				in: func() *int {
					var i int = 1024
					return &i
				}(),
			},
			int(1024),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := structPointerTypeTraitFoo(tt.args.in)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("structPointerTypeTraitFoo() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_structPointerTypeTraitFooWithFunc(t *testing.T) {
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
			structPointerTypeTraitFooWithFunc()
		})
	}
}

func TestSetFoo(t *testing.T) {
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
			SetFoo()
		})
	}
}
