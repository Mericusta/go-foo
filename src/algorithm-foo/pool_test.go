package algorithmfoo

import "testing"

func Test_poolFoo(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test case 1",
			args{c: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			poolFoo(tt.args.c)
		})
	}
}

func Test_poolFooOrigin(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test case 1",
			args{c: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			poolFooOrigin(tt.args.c)
		})
	}
}

func Test_poolFooCompare(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test case 1",
			args{c: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			poolFooCompare(tt.args.c)
		})
	}
}

func Test_poolFooOrigin1(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test case 1",
			args{c: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			poolFooOrigin1(tt.args.c)
		})
	}
}

func Test_poolFooCompare1(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test case 1",
			args{c: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			poolFooCompare1(tt.args.c)
		})
	}
}

func Test_poolFooOrigin2(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test case 1",
			args{c: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			poolFooOrigin2(tt.args.c)
		})
	}
}

func Test_poolFooCompare2(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test case 1",
			args{c: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			poolFooCompare2(tt.args.c)
		})
	}
}

func Test_poolFooOrigin3(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test case 1",
			args{c: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			poolFooOrigin3(tt.args.c)
		})
	}
}

func Test_poolFooCompare3(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test case 1",
			args{c: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			poolFooCompare3(tt.args.c)
		})
	}
}
