package randfoo

import (
	"testing"
	"time"
)

func BenchmarkGetRandSlice(b *testing.B) {
	type args struct {
		seed int64
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{time.Now().UnixNano() % time.Now().Unix()},
			[]int{1, 0, 1, 0},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			GetRandSlice(tt.args.seed)
		}
	}
	t.Logf("[ seed = %v ]\n", tests[0].args.seed)
}

func BenchmarkRandSlice(b *testing.B) {
	type args struct {
		seed      int64
		otherInfo string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			RandSlice(tt.args.seed, tt.args.otherInfo)
		}
	}
}
