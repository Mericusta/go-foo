package benchmarkfoo

import (
	"testing"
	"time"
)

func Benchmark_queueFoo(b *testing.B) {
	type args struct {
		produceCount  int
		producerCount int
		p             PoolDequeue
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{
		// {
		// 	"test case 1",
		// 	time.Second,
		// 	args{
		// 		produceCount:  1,
		// 		producerCount: 1024,
		// 		p:             NewQueue[int](),
		// 	},
		// },
		{
			"test case 2",
			time.Second,
			args{
				produceCount:  1,
				producerCount: 1024,
				p:             NewPoolDequeue(1024),
			},
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			queueFoo(tt.args.produceCount, tt.args.producerCount, tt.args.p)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}
