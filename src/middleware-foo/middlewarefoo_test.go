package middlewarefoo

import "testing"

func TestHandlerMiddlewareFoo(t *testing.T) {
	type args struct {
		ia                         interfaceA
		withUserContext            bool
		withOtherServerUserContext bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				ia:                         &implementA{},
				withUserContext:            false,
				withOtherServerUserContext: false,
			},
		},
		{
			"test case 2",
			args{
				ia:                         &implementA{},
				withUserContext:            true,
				withOtherServerUserContext: false,
			},
		},
		{
			"test case 3",
			args{
				ia:                         &implementA{},
				withUserContext:            false,
				withOtherServerUserContext: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandlerMiddlewareFoo(tt.args.ia, tt.args.withUserContext, tt.args.withOtherServerUserContext)
		})
	}
}
