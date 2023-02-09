package algorithmfoo

import (
	"reflect"
	"testing"
)

func TestPutBackRandom(t *testing.T) {
	type args struct {
		count     int
		rangeFunc func(func(int, int) bool)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PutBackRandom(tt.args.count, tt.args.rangeFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutBackRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPutAwayRandom(t *testing.T) {
	type args struct {
		count     int
		rangeFunc func(func(int, int) bool)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PutAwayRandom(tt.args.count, tt.args.rangeFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutAwayRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}
