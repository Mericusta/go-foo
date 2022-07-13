package httpfoo

import "testing"

func TestRequestExample(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RequestExample(tt.args.index)
		})
	}
}
