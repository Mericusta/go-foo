package redisfoo

import "testing"

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
