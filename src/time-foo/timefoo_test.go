package timefoo

import "testing"

func Test_tickerAndSleep(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{count: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tickerAndSleep(tt.args.count)
		})
	}
}
