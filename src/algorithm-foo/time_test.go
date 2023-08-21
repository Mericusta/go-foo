package algorithmfoo

import (
	"testing"
)

func Test_MiddleOfYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test case 1",
			args{2023},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MiddleOfYear(tt.args.year)
		})
	}
}
