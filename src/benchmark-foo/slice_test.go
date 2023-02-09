package benchmarkfoo

import (
	"reflect"
	"testing"
)

func Test_removeDuplication_map(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplication_map(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplication_map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplication_sort(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplication_sort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplication_sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swap(t *testing.T) {
	type args struct {
		arr []string
		a   int
		b   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swap(tt.args.arr, tt.args.a, tt.args.b)
		})
	}
}

func Test_simple_removeDuplication_map(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simple_removeDuplication_map(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simple_removeDuplication_map() = %v, want %v", got, tt.want)
			}
		})
	}
}
