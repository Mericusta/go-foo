package algorithmfoo

import "testing"

func Test_my_fib(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			1,
			1,
		},
		{
			"test case 5",
			5,
			8,
		},
		{
			"test case 10",
			10,
			89,
		},
		{
			"test case 40",
			40,
			165580141,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := my_fib(tt.n)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("my_fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_my_fib_cycle(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			1,
			1,
		},
		{
			"test case 5",
			5,
			8,
		},
		{
			"test case 10",
			10,
			89,
		},
		{
			"test case 40",
			40,
			165580141,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := my_fib_cycle(tt.n)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("my_fib_cycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
