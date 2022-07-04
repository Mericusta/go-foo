package mapfoo

import "testing"

func TestMapCapacityFoo(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1: c 4",
			args{c: 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MapCapacityFoo(tt.args.c)
		})
	}
}
