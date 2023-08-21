package benchmarkfoo

import "testing"

func Test_queueFoo(t *testing.T) {
	type args struct {
		produceCount  int
		producerCount int
		p             PoolDequeue
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				produceCount:  1,
				producerCount: 128,
				p:             NewQueue[int](),
			},
		},
		{
			"test case 2",
			args{
				produceCount:  1,
				producerCount: 128,
				p:             NewPoolDequeue(128),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queueFoo(tt.args.produceCount, tt.args.producerCount, tt.args.p)
		})
	}
}
