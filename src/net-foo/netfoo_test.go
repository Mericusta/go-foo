package netfoo

import (
	"testing"
)

func TestCloseConnectorFoo(t *testing.T) {
	type args struct {
		closedBy int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		// {
		// 	"test case 1, connection closed by server",
		// 	args{closedBy: 1},
		// },
		// {
		// 	"test case 2, connection closed by client",
		// 	args{closedBy: 2},
		// },
		{
			"test case 2, connection closed by client",
			args{closedBy: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CloseConnectorFoo(tt.args.closedBy)
		})
	}
}

func TestCloseAndReconnectFoo(t *testing.T) {
	type args struct {
		reconnectCount int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{reconnectCount: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CloseAndReconnectFoo(tt.args.reconnectCount)
		})
	}
}

func Test_tlvFoo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"test case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tlvFoo()
		})
	}
}

func Test_tlvNetBufferFoo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"test case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tlvNetBufferFoo()
		})
	}
}
