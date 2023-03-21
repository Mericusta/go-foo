package algorithmfoo

import (
	"testing"
	"time"
)

// 4096 Benchmark_poolFooOrigin-12    	       1	1104006900 ns/op	70809544 B/op	    4517 allocs/op
// 2048 Benchmark_poolFooOrigin-12    	       3	 417412100 ns/op	17779664 B/op	    2254 allocs/op
// 2048 no force GC Benchmark_poolFooOrigin-12    	      10	 103718360 ns/op	17760373 B/op	    2050 allocs/op
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
// 2048 no force GC Benchmark_poolFooCompare-12    	       9	 127957522 ns/op	18219051 B/op	    4278 allocs/op
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

// 2048 Benchmark_poolFooOrigin2-12    	       3	 388019667 ns/op	17781616 B/op	    2274 allocs/op
// Benchmark_poolFooOrigin2-12    	    4945	    259568 ns/op	  163841 B/op	       1 allocs/op
func Benchmark_poolFooOrigin2(b *testing.B) {
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
			args{c: 2048},
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			poolFooOrigin2(tt.args.c)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

// 2048 Benchmark_poolFooOrigin2-12    	       3	 388019667 ns/op	17781616 B/op	    2274 allocs/op
// 2048 Benchmark_poolFooCompare2-12    	   2	 882093400 ns/op	66593364 B/op	   38019 allocs/op
// Benchmark_poolFooOrigin2-12    	    4945	    259568 ns/op	  163841 B/op	       1 allocs/op
// Benchmark_poolFooOrigin2-12    	    7189	    174449 ns/op	  245763 B/op	       1 allocs/op
// Benchmark_poolFooCompare2-12    	   10000	    116517 ns/op	    1627 B/op	       3 allocs/op

// Benchmark_poolFooOrigin2-12    	    2592	    473660 ns/op	  245779 B/op	       1 allocs/op
// Benchmark_poolFooCompare2-12    	    2744	    440724 ns/op	    3173 B/op	       6 allocs/op

// Benchmark_poolFooOrigin2-12    	      24	  51965258 ns/op	17760324 B/op	    2050 allocs/op
// Benchmark_poolFooCompare2-12    	      15	  75842033 ns/op	18275462 B/op	    4297 allocs/op
func Benchmark_poolFooCompare2(b *testing.B) {
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
			args{c: 2048},
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			poolFooCompare2(tt.args.c)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

// Benchmark_poolFooOrigin3-12    	       6	 167846017 ns/op	222935582 B/op	 2100231 allocs/op
func Benchmark_poolFooOrigin3(b *testing.B) {
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
			args{c: 2048},
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			poolFooOrigin3(tt.args.c)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

func Benchmark_poolFooCompare3(b *testing.B) {
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
			args{c: 2048},
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			poolFooCompare3(tt.args.c)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}
