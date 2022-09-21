package syncfoo

import (
	"testing"
	"time"
)

func BenchmarkMutexLockerPerformanceOnBlockingGoroutine(b *testing.B) {
	type args struct {
		gCount int
		d      time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				gCount: 100,
				d:      time.Millisecond * 10,
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			MutexLockerPerformanceOnBlockingGoroutine(tt.args.gCount, tt.args.d)
		}
	}
}

func BenchmarkMutexLockerPerformanceOnChannelReceiver(b *testing.B) {
	type args struct {
		gCount         int
		tickerDuration time.Duration
		tickerMax      int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				gCount:         100,
				tickerDuration: time.Millisecond * 10,
				tickerMax:      1000,
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			MutexLockerPerformanceOnChannelReceiver(tt.args.gCount, tt.args.tickerDuration, tt.args.tickerMax)
		}
	}
}

func BenchmarkMutexLockerPerformanceOnHttpRequest(b *testing.B) {
	type args struct {
		gCount int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{gCount: 100},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			MutexLockerPerformanceOnHttpRequest(tt.args.gCount)
		}
	}
}

func BenchmarkMutexLockerPerformanceOnLoadCacheFromRedis(b *testing.B) {
	type args struct {
		gCount int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int32
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{gCount: 100},
			"Hello Spin Key",
			100,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			MutexLockerPerformanceOnLoadCacheFromRedis(tt.args.gCount)
		}
	}
}

func BenchmarkMutexLockerPerformanceOnLocalOperation(b *testing.B) {
	type args struct {
		gCount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{gCount: 100},
			100,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			MutexLockerPerformanceOnLocalOperation(tt.args.gCount)
		}
	}
}

func BenchmarkRedisV8CacheOncePerformanceOnLoadCacheFromRedis(b *testing.B) {
	type args struct {
		gCount int
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{gCount: 100},
			100,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			RedisV8CacheOncePerformanceOnLoadCacheFromRedis(tt.args.gCount)
		}
	}
}

func BenchmarkRedisV8CachePerformanceOnLoadCacheFromRedis(b *testing.B) {
	type args struct {
		gCount int
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{gCount: 100},
			100,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			RedisV8CachePerformanceOnLoadCacheFromRedis(tt.args.gCount)
		}
	}
}

func BenchmarkSequentialGroupOnLocalOperation(b *testing.B) {
	type args struct {
		gCount     int
		groupCount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	"10w 1, 8578	    137918 ns/op	      80 B/op	       4 allocs/op",
		// 	args{
		// 		gCount:     100000,
		// 		groupCount: 1,
		// 	},
		// 	100000,
		// },
		// {
		// 	"10w 10, 7275	    190213 ns/op	     600 B/op	      22 allocs/op",
		// 	args{
		// 		gCount:     100000,
		// 		groupCount: 10,
		// 	},
		// 	100000,
		// },
		// {
		// 	"10w 100, 5308	    224366 ns/op	    5657 B/op	     202 allocs/op",
		// 	args{
		// 		gCount:     100000,
		// 		groupCount: 100,
		// 	},
		// 	100000,
		// },
		// {
		// 	"10w 1000, 2942	    399963 ns/op	   56180 B/op	    2003 allocs/op",
		// 	args{
		// 		gCount:     100000,
		// 		groupCount: 1000,
		// 	},
		// 	100000,
		// },
		// {
		// 	"10w 10000, 385	   3167677 ns/op	  560581 B/op	   20004 allocs/op",
		// 	args{
		// 		gCount:     100000,
		// 		groupCount: 10000,
		// 	},
		// 	100000,
		// },
		// {
		// 	"10w 100000, 39	  29077362 ns/op	 5607544 B/op	  200035 allocs/op",
		// 	args{
		// 		gCount:     100000,
		// 		groupCount: 100000,
		// 	},
		// 	100000,
		// },
		// {
		// 	"100w 1, 829	   1427389 ns/op	      80 B/op	       4 allocs/op",
		// 	args{
		// 		gCount:     1000000,
		// 		groupCount: 1,
		// 	},
		// 	1000000,
		// },
		// {
		// 	"100w 10, 718	   1414013 ns/op	     625 B/op	      22 allocs/op",
		// 	args{
		// 		gCount:     1000000,
		// 		groupCount: 10,
		// 	},
		// 	1000000,
		// },
		// {
		// 	"100w 100, 771	   1548260 ns/op	    5680 B/op	     202 allocs/op",
		// 	args{
		// 		gCount:     1000000,
		// 		groupCount: 100,
		// 	},
		// 	1000000,
		// },
		// {
		// 	"100w 1000, 600	   2022531 ns/op	   56689 B/op	    2008 allocs/op",
		// 	args{
		// 		gCount:     1000000,
		// 		groupCount: 1000,
		// 	},
		// 	1000000,
		// },
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			SequentialGroupOnLocalOperation(tt.args.gCount, tt.args.groupCount)
		}
	}
}

func BenchmarkSingleFlightPerformanceOnLoadCacheFromRedis(b *testing.B) {
	type args struct {
		gCount int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int32
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{gCount: 2},
			"Hello Spin Key",
			2,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			SingleFlightPerformanceOnLoadCacheFromRedis(tt.args.gCount)
		}
	}
}

func BenchmarkSpinLockerPerformanceOnBlockingGoroutine(b *testing.B) {
	type args struct {
		gCount int
		d      time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				gCount: 100,
				d:      time.Millisecond * 10,
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			SpinLockerPerformanceOnBlockingGoroutine(tt.args.gCount, tt.args.d)
		}
	}
}

func BenchmarkSpinLockerPerformanceOnChannelReceiver(b *testing.B) {
	type args struct {
		gCount         int
		tickerDuration time.Duration
		tickerMax      int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				gCount:         100,
				tickerDuration: time.Millisecond * 10,
				tickerMax:      1000,
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			SpinLockerPerformanceOnChannelReceiver(tt.args.gCount, tt.args.tickerDuration, tt.args.tickerMax)
		}
	}
}

func BenchmarkSpinLockerPerformanceOnHttpRequest(b *testing.B) {
	type args struct {
		gCount int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{gCount: 100},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			SpinLockerPerformanceOnHttpRequest(tt.args.gCount)
		}
	}
}

func BenchmarkSpinLockerPerformanceOnLoadCacheFromRedis(b *testing.B) {
	type args struct {
		gCount int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int32
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{gCount: 100},
			"Hello Spin Key",
			100,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			SpinLockerPerformanceOnLoadCacheFromRedis(tt.args.gCount)
		}
	}
}

func BenchmarkSpinLockerPerformanceOnLocalOperation(b *testing.B) {
	type args struct {
		gCount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{gCount: 100},
			100,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			SpinLockerPerformanceOnLocalOperation(tt.args.gCount)
		}
	}
}

func BenchmarkTidwallSpinLockerPerformanceOnBlockingGoroutine(b *testing.B) {
	type args struct {
		gCount int
		d      time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				gCount: 100,
				d:      time.Millisecond * 10,
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			TidwallSpinLockerPerformanceOnBlockingGoroutine(tt.args.gCount, tt.args.d)
		}
	}
}

func BenchmarkTidwallSpinLockerPerformanceOnChannelReceiver(b *testing.B) {
	type args struct {
		gCount         int
		tickerDuration time.Duration
		tickerMax      int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				gCount:         100,
				tickerDuration: time.Millisecond * 10,
				tickerMax:      1000,
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			TidwallSpinLockerPerformanceOnChannelReceiver(tt.args.gCount, tt.args.tickerDuration, tt.args.tickerMax)
		}
	}
}

func BenchmarkTidwallSpinLockerPerformanceOnHttpRequest(b *testing.B) {
	type args struct {
		gCount int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{gCount: 100},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			TidwallSpinLockerPerformanceOnHttpRequest(tt.args.gCount)
		}
	}
}

func BenchmarkTidwallSpinLockerPerformanceOnLoadCacheFromRedis(b *testing.B) {
	type args struct {
		gCount int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int32
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{gCount: 100},
			"Hello Spin Key",
			100,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			TidwallSpinLockerPerformanceOnLoadCacheFromRedis(tt.args.gCount)
		}
	}
}

func BenchmarkTidwallSpinLockerPerformanceOnLocalOperation(b *testing.B) {
	type args struct {
		gCount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{gCount: 100},
			100,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			TidwallSpinLockerPerformanceOnLocalOperation(tt.args.gCount)
		}
	}
}

func BenchmarkWaitGroupCallFunctionWillCaptureWhenWait(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			WaitGroupCallFunctionWillCaptureWhenWait()
		}
	}
}
