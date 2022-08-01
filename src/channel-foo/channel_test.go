package channelfoo

import (
	"testing"
)

func TestGoroutineExitThenCloseChannel(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GoroutineExitThenCloseChannel()
		})
	}
}

func TestListenerBlockedChannel(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListenerBlockedChannel()
		})
	}
}

func TestGoroutineExitThenCloseChannelSimpleCase(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GoroutineExitThenCloseChannelSimpleCase()
		})
	}
}

func TestGoroutineOutputOrder(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		// TODO: Add test cases.
		{
			"test case",
			0,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GoroutineOutputOrder()
			if got != tt.want {
				t.Errorf("GoroutineOutputOrder() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GoroutineOutputOrder() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGoroutineOutputOrder2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GoroutineOutputOrder2()
		})
	}
}

func TestGoChannelBlock(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GoChannelBlock()
		})
	}
}
