package builtinfoo

import "testing"

func TestGoroutinePassObjectPointerFoo(t *testing.T) {
	type args struct {
		generatePointer bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{generatePointer: false},
		},
		{
			"test case 1",
			args{generatePointer: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GoroutinePassObjectPointerFoo(tt.args.generatePointer)
		})
	}
}
