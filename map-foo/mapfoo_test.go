package mapfoo

import (
	"testing"
)

func TestMapCapacityFoo(t *testing.T) {
	type args struct {
		count    int
		capacity int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case",
			args{count: 8, capacity: 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MapCapacityFoo(tt.args.count, tt.args.capacity)
		})
	}
}

func TestStructMapKeyFoo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StructMapKeyFoo()
		})
	}
}
