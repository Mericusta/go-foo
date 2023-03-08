package algorithmfoo

import (
	"testing"
	"time"
)

// 4096 Benchmark_poolFooOrigin-12    	       1	1104006900 ns/op	70809544 B/op	    4517 allocs/op
// 2048 Benchmark_poolFooOrigin-12    	       3	 417412100 ns/op	17779664 B/op	    2254 allocs/op
func Benchmark_poolFooOrigin(b *testing.B) {
	type args struct {
		c int
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{
		{
			"test case 1",
			time.Second * 2,
			args{c: 2048}, // (1+2048)*2048/2
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			poolFooOrigin(tt.args.c)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %vms\n", tt.limit, time.Duration(float64(b.Elapsed())/float64(b.N)).Milliseconds())
		}
	}
}

// 4096 Benchmark_poolFooCompare-12    	       1	2878158700 ns/op	258073392 B/op	   84011 allocs/op
// 2048 Benchmark_poolFooCompare-12    	       2	 871308350 ns/op	66579036 B/op	   37956 allocs/op
func Benchmark_poolFooCompare(b *testing.B) {
	type args struct {
		c int
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{
		{
			"test case 1",
			time.Second * 3,
			args{c: 2048}, // (1+2048)*2048/2
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			poolFooCompare(tt.args.c)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

// 2048 Benchmark_poolFooOrigin1-12    	       2	 594940700 ns/op	17990928 B/op	    2240 allocs/op
func Benchmark_poolFooOrigin1(b *testing.B) {
	type args struct {
		c int
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{
		{
			"test case 1",
			time.Second * 3,
			args{c: 2048}, // (1+2048)*2048/2
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			poolFooOrigin1(tt.args.c)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

// Benchmark_poolFooCompare1-12    	       1	2424009500 ns/op	264184104 B/op	 6346883 allocs/op
func Benchmark_poolFooCompare1(b *testing.B) {
	type args struct {
		c int
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{
		{
			"test case 1",
			time.Second * 3,
			args{c: 2048}, // (1+2048)*2048/2
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			poolFooCompare1(tt.args.c)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}
