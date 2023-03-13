package packagefoo

import (
	"reflect"
	"testing"
)

func Test_stringsTrimFoo(t *testing.T) {
	type args struct {
		s      string
		cutset string
	}
	tests := []struct {
		name  string
		args  args
		want0 string
	}{
		{
			"test case 1",
			args{"strings_foo.go", ".go"},
			"strings_foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := stringsTrimFoo(tt.args.s, tt.args.cutset)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("stringsTrimFoo() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}
