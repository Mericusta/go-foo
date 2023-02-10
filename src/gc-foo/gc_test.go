package gcfoo

import (
	"testing"
)

func TestForceGCPointerSlice(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{c: 1e9},
		},
		{
			"test case 2",
			args{c: 1024},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ForceGCPointerSlice(tt.args.c)
		})
	}
}

func TestForceGCNonPointerSlice(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{c: 1e9},
		},
		{
			"test case 2",
			args{c: 1024},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ForceGCNonPointerSlice(tt.args.c)
		})
	}
}

func TestForceGCPointerSliceInOSHeap(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{c: 1e9},
		},
		{
			"test case 2",
			args{c: 1024},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ForceGCPointerSliceInOSHeap(tt.args.c)
		})
	}
}

func TestForceGCNoNeedReleaseString(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{c: 1 << 24},
		},
		{
			"test case 2",
			args{c: 1024},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ForceGCNoNeedReleaseString(tt.args.c)
		})
	}
}

func TestAvoidForceGCNoNeedReleaseString(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{c: 1 << 24},
		},
		{
			"test case 2",
			args{c: 1024},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AvoidForceGCNoNeedReleaseString(tt.args.c)
		})
	}
}
