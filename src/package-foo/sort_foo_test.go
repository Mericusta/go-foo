package packagefoo

import (
	"reflect"
	"testing"
)

func TestSortFoo(t *testing.T) {
	type args struct {
		s       []int
		infos   []*info
		reverse bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				s: []int{2, 1, 3},
				infos: []*info{
					{ID: 2, Value: 2},
					{ID: 1, Value: 1},
					{ID: 3, Value: 3},
				},
				reverse: false,
			},
			[]int{1, 2, 3},
		},
		{
			"test case 2",
			args{
				s: []int{2, 1, 3},
				infos: []*info{
					{ID: 2, Value: 2},
					{ID: 1, Value: 1},
					{ID: 3, Value: 3},
				},
				reverse: true,
			},
			[]int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortFoo(tt.args.s, tt.args.infos, tt.args.reverse); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortFoo() = %v, want %v", got, tt.want)
			}
		})
	}
}
