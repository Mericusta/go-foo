package algorithmfoo

import "testing"

func TestMonteCarloMethod_Estimating_PI(t *testing.T) {
	type args struct {
		points int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{points: 100000000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MonteCarloMethod_Estimating_PI(tt.args.points)
		})
	}
}
