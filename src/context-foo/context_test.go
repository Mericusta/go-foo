package contextfoo

import "testing"

func TestTimeoutContextFoo(t *testing.T) {
	type args struct {
		timeoutSeconds  int
		businessSeconds int
		businessPanic   bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				timeoutSeconds:  2,
				businessSeconds: 10,
				businessPanic:   false,
			},
		},
		{
			"test case 2",
			args{
				timeoutSeconds:  2,
				businessSeconds: 10,
				businessPanic:   true,
			},
		},
		// {
		// 	"test case 2",
		// 	args{
		// 		timeoutSeconds:  3,
		// 		businessSeconds: 4,
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TimeoutContextFoo(tt.args.timeoutSeconds, tt.args.businessSeconds, tt.args.businessPanic)
		})
	}
}
