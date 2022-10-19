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

func TestGetFromMapAsTypeEmptyValueFoo(t *testing.T) {
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
			GetFromMapAsTypeEmptyValueFoo()
		})
	}
}

func TestReadConcurrently(t *testing.T) {
	type args struct {
		c int
		s int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{c: 3000, s: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadConcurrently(tt.args.c, tt.args.s)
		})
	}
}
