package builtinfoo

import (
	"reflect"
	"testing"
)

func Test_DeepEqualFoo_1af009229e742cbb0d358bba9b48aaea(t *testing.T) {
	type args struct {
		t any
		v int64
	}
	tests := []struct {
		name  string
		args  args
		want0 bool
	}{
		{
			"test case 1",
			args{
				t: nil,
				v: 3,
			},
			false,
		},
		{
			"test case 2",
			args{
				t: nil,
				v: 0,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := DeepEqualFoo(tt.args.t, tt.args.v)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("DeepEqualFoo() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_DeepEqualFoo_b5b621322159eff32ebce5a73d59d21e(t *testing.T) {
	type args struct {
		t any
		v any
	}
	tests := []struct {
		name  string
		args  args
		want0 bool
	}{
		{
			"test case 1",
			args{
				t: nil,
				v: 3,
			},
			false,
		},
		{
			"test case 2",
			args{
				t: nil,
				v: 0,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := DeepEqualFoo(tt.args.t, tt.args.v)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("DeepEqualFoo() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_DeepEqualFoo_632759eb51780f187fb028d13a5b9115(t *testing.T) {
	type args struct {
		t interface{}
		v interface{}
	}
	tests := []struct {
		name  string
		args  args
		want0 bool
	}{
		{
			"test case 1",
			args{
				t: nil,
				v: 3,
			},
			false,
		},
		{
			"test case 2",
			args{
				t: nil,
				v: 0,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := DeepEqualFoo(tt.args.t, tt.args.v)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("DeepEqualFoo() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}
