package middlewarefoo

import "testing"

func TestHandlerMiddlewareFoo(t *testing.T) {
	type args struct {
		ia interfaceA
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{ia: newInterfaceA()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandlerMiddlewareFoo(tt.args.ia)
		})
	}
}
