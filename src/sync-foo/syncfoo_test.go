package syncfoo

import (
	"testing"
	"time"
)

func TestWaitGroupCallFunctionWillCaptureWhenWait(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WaitGroupCallFunctionWillCaptureWhenWait()
		})
	}
}

func TestSpinLockerPerformanceOnLocalOperation(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpinLockerPerformanceOnLocalOperation(tt.args.gCount); got != tt.want {
				t.Errorf("SpinLockerPerformanceOnLocalOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpinLockerPerformanceOnLoadCacheFromRedis(t *testing.T) {
	type args struct {
		gCount int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{gCount: 100},
			redisCacheValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpinLockerPerformanceOnLoadCacheFromRedis(tt.args.gCount); got != tt.want {
				t.Errorf("SpinLockerPerformanceOnLoadCacheFromRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMutexLockerPerformanceOnLocalOperation(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MutexLockerPerformanceOnLocalOperation(tt.args.gCount); got != tt.want {
				t.Errorf("MutexLockerPerformanceOnLocalOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMutexLockerPerformanceOnLoadCacheFromRedis(t *testing.T) {
	type args struct {
		gCount int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{gCount: 100},
			redisCacheValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MutexLockerPerformanceOnLoadCacheFromRedis(tt.args.gCount); got != tt.want {
				t.Errorf("MutexLockerPerformanceOnLoadCacheFromRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpinLockerPerformanceOnBlockingGoroutine(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SpinLockerPerformanceOnBlockingGoroutine(tt.args.gCount, tt.args.d)
		})
	}
}

func TestMutexLockerPerformanceOnBlockingGoroutine(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MutexLockerPerformanceOnBlockingGoroutine(tt.args.gCount, tt.args.d)
		})
	}
}

func TestSpinLockerPerformanceOnChannelReceiver(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SpinLockerPerformanceOnChannelReceiver(tt.args.gCount, tt.args.tickerDuration, tt.args.tickerMax)
		})
	}
}

func TestMutexLockerPerformanceOnChannelReceiver(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MutexLockerPerformanceOnChannelReceiver(tt.args.gCount, tt.args.tickerDuration, tt.args.tickerMax)
		})
	}
}

func TestSpinLockerPerformanceOnHttpRequest(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SpinLockerPerformanceOnHttpRequest(tt.args.gCount)
		})
	}
}

func TestMutexLockerPerformanceOnHttpRequest(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MutexLockerPerformanceOnHttpRequest(tt.args.gCount)
		})
	}
}

func TestTidwallSpinLockerPerformanceOnLocalOperation(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TidwallSpinLockerPerformanceOnLocalOperation(tt.args.gCount); got != tt.want {
				t.Errorf("TidwallSpinLockerPerformanceOnLocalOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTidwallSpinLockerPerformanceOnLoadCacheFromRedis(t *testing.T) {
	type args struct {
		gCount int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{gCount: 100},
			redisCacheValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TidwallSpinLockerPerformanceOnLoadCacheFromRedis(tt.args.gCount); got != tt.want {
				t.Errorf("TidwallSpinLockerPerformanceOnLoadCacheFromRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTidwallSpinLockerPerformanceOnBlockingGoroutine(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TidwallSpinLockerPerformanceOnBlockingGoroutine(tt.args.gCount, tt.args.d)
		})
	}
}

func TestTidwallSpinLockerPerformanceOnChannelReceiver(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TidwallSpinLockerPerformanceOnChannelReceiver(tt.args.gCount, tt.args.tickerDuration, tt.args.tickerMax)
		})
	}
}

func TestTidwallSpinLockerPerformanceOnHttpRequest(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TidwallSpinLockerPerformanceOnHttpRequest(tt.args.gCount)
		})
	}
}
