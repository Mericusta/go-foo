package randfoo

import (
	"reflect"
	"testing"
	"time"
)

func TestGetRandSlice(t *testing.T) {
	type args struct {
		seed int64
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{time.Now().UnixNano() % time.Now().Unix()},
			[]int{1, 0, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandSlice(tt.args.seed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRandSlice() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Logf("[ seed = %v ]\n", tests[0].args.seed)
}

func TestRandSlice(t *testing.T) {
	type args struct {
		seed      int64
		otherInfo string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RandSlice(tt.args.seed, tt.args.otherInfo)
		})
	}
}
