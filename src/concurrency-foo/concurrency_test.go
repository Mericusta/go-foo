package concurrencyfoo

import (
	"testing"
)

func TestGoroutineCommunicateByBufferChannelWithLittleStructFoo(t *testing.T) {
	type args struct {
		senderCount int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{senderCount: 10000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GoroutineCommunicateByBufferChannelWithLittleStructFoo(tt.args.senderCount)
		})
	}
}

func TestSyncFoolWithSliceFoo(t *testing.T) {
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
			SyncFoolWithSliceFoo()
		})
	}
}

func Test_atomicFoo(t *testing.T) {
	type args struct {
		withAtomicWrite bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{withAtomicWrite: false},
		},
		// {
		// 	"test case 1",
		// 	args{true},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			atomicFoo(tt.args.withAtomicWrite)
		})
	}
}
