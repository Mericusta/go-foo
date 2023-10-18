package redisfoo

import (
	"testing"
)

func Test_connectTest(t *testing.T) {
	type args struct {
		url      string
		password string
		DB       int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				url:      "192.168.2.203:6379",
				password: "",
				DB:       0,
			},
			"PONG",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connectTest(tt.args.url, tt.args.password, tt.args.DB); got != tt.want {
				t.Errorf("connectTest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_zaddFoo(t *testing.T) {
	type args struct {
		url      string
		password string
		DB       int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				url:      "192.168.2.147:6379",
				password: "",
				DB:       1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zaddFoo(tt.args.url, tt.args.password, tt.args.DB)
		})
	}
}

func Test_zrankFoo(t *testing.T) {
	type args struct {
		url      string
		password string
		DB       int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				url:      "192.168.2.147:6379",
				password: "",
				DB:       1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zrankFoo(tt.args.url, tt.args.password, tt.args.DB)
		})
	}
}

func Test_zrevrankFoo(t *testing.T) {
	type args struct {
		url      string
		password string
		DB       int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				url:      "192.168.2.147:6379",
				password: "",
				DB:       1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zrevrankFoo(tt.args.url, tt.args.password, tt.args.DB)
		})
	}
}

func Test_getFoo(t *testing.T) {
	type args struct {
		url      string
		password string
		DB       int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				url:      "192.168.2.147:6379",
				password: "",
				DB:       1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getFoo(tt.args.url, tt.args.password, tt.args.DB)
		})
	}
}

func Test_hsetFoo(t *testing.T) {
	type args struct {
		url      string
		password string
		DB       int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				url:      "192.168.2.147:6379",
				password: "",
				DB:       1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsetFoo(tt.args.url, tt.args.password, tt.args.DB)
		})
	}
}

func Test_distributedLockerFoo(t *testing.T) {
	type args struct {
		url      string
		password string
		DB       int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				url:      "192.168.2.147:6379",
				password: "",
				DB:       1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			distributedLockerFoo(tt.args.url, tt.args.password, tt.args.DB)
		})
	}
}

func Test_hgetallFoo(t *testing.T) {
	type args struct {
		url      string
		password string
		DB       int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				url:      "192.168.2.147:6379",
				password: "",
				DB:       1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hgetallFoo(tt.args.url, tt.args.password, tt.args.DB)
		})
	}
}
