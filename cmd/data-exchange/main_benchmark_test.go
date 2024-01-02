package main

import (
	data_exchange "go-foo/cmd/data-exchange/pb"
	"testing"
	"time"
)

// Benchmark_ProtoMarshalFoo_17ed74d8fa7607506f6248e12df7da56-12    	  206768	      5970 ns/op	    1824 B/op	      91 allocs/op
func Benchmark_ProtoMarshalFoo_17ed74d8fa7607506f6248e12df7da56(b *testing.B) {
	type args struct {
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{
		{
			"test case 1",
			100 * time.Millisecond,
			args{v: newRobotsFightData()},
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = ProtoMarshalFoo(tt.args.v)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

// Benchmark_ProtoUnmarshalFoo_17ed74d8fa7607506f6248e12df7da56-12    	16844019	        71.97 ns/op	       0 B/op	       0 allocs/op
func Benchmark_ProtoUnmarshalFoo_17ed74d8fa7607506f6248e12df7da56(b *testing.B) {
	type args struct {
		b []byte
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{
		{
			"test case 1",
			100 * time.Millisecond,
			args{v: newRobotsFightData()},
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = ProtoUnmarshalFoo(tt.args.b, tt.args.v)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

// Benchmark_JsonMarshalFoo_17ed74d8fa7607506f6248e12df7da56-12    	  161761	      6850 ns/op	    2889 B/op	      61 allocs/op
func Benchmark_JsonMarshalFoo_17ed74d8fa7607506f6248e12df7da56(b *testing.B) {
	type args struct {
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{
		{
			"test case 1",
			100 * time.Millisecond,
			args{v: newRobotsFightData()},
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = JsonMarshalFoo(tt.args.v)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

// Benchmark_JsonUnmarshalFoo_17ed74d8fa7607506f6248e12df7da56-12    	12516962	        96.22 ns/op	     168 B/op	       2 allocs/op
func Benchmark_JsonUnmarshalFoo_17ed74d8fa7607506f6248e12df7da56(b *testing.B) {
	type args struct {
		b []byte
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{
		{
			"test case 1",
			100 * time.Millisecond,
			args{v: newRobotsFightData()},
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = JsonUnmarshalFoo(tt.args.b, tt.args.v)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

// Benchmark_GobMarshalFoo_17ed74d8fa7607506f6248e12df7da56-12    	  135976	      9409 ns/op	    2784 B/op	      91 allocs/op
func Benchmark_GobMarshalFoo_17ed74d8fa7607506f6248e12df7da56(b *testing.B) {
	type args struct {
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{
		{
			"test case 1",
			100 * time.Millisecond,
			args{v: newRobotsFightData()},
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = GobMarshalFoo(tt.args.v)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

// Benchmark_GobUnmarshalFoo_17ed74d8fa7607506f6248e12df7da56-12    	 4551505	       263.4 ns/op	     336 B/op	       6 allocs/op
func Benchmark_GobUnmarshalFoo_17ed74d8fa7607506f6248e12df7da56(b *testing.B) {
	type args struct {
		b []byte
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{
		{
			"test case 1",
			100 * time.Millisecond,
			args{v: newRobotsFightData()},
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = GobUnmarshalFoo(tt.args.b, tt.args.v)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

// Benchmark_ProtoFoo_17ed74d8fa7607506f6248e12df7da56-12    	   87394	     13127 ns/op	    3544 B/op	     160 allocs/op
func Benchmark_ProtoFoo_17ed74d8fa7607506f6248e12df7da56(b *testing.B) {
	type args struct {
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{
		{
			"test case 1",
			100 * time.Millisecond,
			args{v: newRobotsFightData()},
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = ProtoFoo(tt.args.v)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

// Benchmark_JsonFoo_17ed74d8fa7607506f6248e12df7da56-12    	   48186	     25651 ns/op	    3290 B/op	      93 allocs/op
func Benchmark_JsonFoo_17ed74d8fa7607506f6248e12df7da56(b *testing.B) {
	type args struct {
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{
		{
			"test case 1",
			100 * time.Millisecond,
			args{v: newRobotsFightData()},
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = JsonFoo(tt.args.v)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}

// Benchmark_GobFoo_17ed74d8fa7607506f6248e12df7da56-12    	   28282	     43137 ns/op	   13292 B/op	     378 allocs/op
func Benchmark_GobFoo_17ed74d8fa7607506f6248e12df7da56(b *testing.B) {
	type args struct {
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		limit time.Duration
		args  args
	}{
		{
			"test case 1",
			100 * time.Millisecond,
			args{v: newRobotsFightData()},
		},
	}
	for _, tt := range tests {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = GobFoo(tt.args.v)
		}
		b.StopTimer()
		if b.Elapsed() > tt.limit*time.Duration(b.N) {
			b.Fatalf("overtime limit %v, got %.2f\n", tt.limit, float64(b.Elapsed())/float64(b.N))
		}
	}
}
