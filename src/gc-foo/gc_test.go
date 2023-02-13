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

func TestForceGCNoNeedReleaseStringSlice(t *testing.T) {
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
			ForceGCNoNeedReleaseStringSlice(tt.args.c)
		})
	}
}

func TestAvoidForceGCNoNeedReleaseStringSlice(t *testing.T) {
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
			AvoidForceGCNoNeedReleaseStringSlice(tt.args.c)
		})
	}
}

func TestForceGCNoNeedReleaseStringMap(t *testing.T) {
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
			ForceGCNoNeedReleaseStringMap(tt.args.c)
		})
	}
}

func TestAvoidForceGCNoNeedReleaseStringMap(t *testing.T) {
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
			AvoidForceGCNoNeedReleaseStringMap(tt.args.c)
		})
	}
}
