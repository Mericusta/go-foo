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
