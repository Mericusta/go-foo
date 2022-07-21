package builtinfoo

import (
	"reflect"
	"testing"
)

func TestCopyStringSliceFromStringsSplit(t *testing.T) {
	type args struct {
		c       int
		s       string
		lenCase string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			"test case 1, equal",
			args{c: 3, s: "a1,b2,c3", lenCase: "=="},
			[]string{"a1", "b2", "c3"},
		},
		{
			"test case 1, equal",
			args{c: 3, s: "a1,b2,c3", lenCase: "<"},
			[]string{"a1", "b2"},
		},
		{
			"test case 1, equal",
			args{c: 3, s: "a1,b2,c3", lenCase: ">"},
			[]string{"a1", "b2", "c3", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CopyStringSliceFromStringsSplit(tt.args.c, tt.args.s, tt.args.lenCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopyStringSliceFromStringsSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
