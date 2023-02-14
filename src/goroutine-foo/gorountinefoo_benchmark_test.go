package goroutinefoo

import (
	"math"
	"testing"
)

func BenchmarkAllMIsWorking(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			AllMIsWorking()
		}
	}
}

func BenchmarkGoroutinePanicAndRecoverFoo(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			GoroutinePanicAndRecoverFoo()
		}
	}
}

func BenchmarkOpenSoMuchGoRoutine(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			OpenSoMuchGoRoutine()
		}
	}
}

func BenchmarkRecoverAtGoroutine(b *testing.B) {
	type args struct {
		gCount       int
		hCount       int
		mod          int
		concurrently bool
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 int32
	}{
		// TODO: Add test cases.
		// {
		// 	"test case 1, panic probability 1/7, 128034433 ns/op",
		// 	args{
		// 		gCount:       1234,
		// 		hCount:       4321,
		// 		concurrently: true,
		// 		mod:          7,
		// 	},
		// 	func() int32 {
		// 		c := int32(0)
		// 		for i := 0; i != 4321; i++ {
		// 			if i%7 == 0 {
		// 				c++
		// 			}
		// 		}
		// 		return 1234*4321 - 1234*(c-1)
		// 	}(),
		// 	func() int32 {
		// 		c := int32(0)
		// 		for i := 0; i != 4321; i++ {
		// 			if i%7 == 0 {
		// 				c++
		// 			}
		// 		}
		// 		return 1234 * (c - 1)
		// 	}(),
		// },
		// {
		// 	"test case 2, panic probability 1/1024, 100626700 ns/op",
		// 	args{
		// 		gCount:       1234,
		// 		hCount:       4321,
		// 		concurrently: true,
		// 		mod:          1024,
		// 	},
		// 	func() int32 {
		// 		c := int32(0)
		// 		for i := 0; i != 4321; i++ {
		// 			if i%1024 == 0 {
		// 				c++
		// 			}
		// 		}
		// 		return 1234*4321 - 1234*(c-1)
		// 	}(),
		// 	func() int32 {
		// 		c := int32(0)
		// 		for i := 0; i != 4321; i++ {
		// 			if i%1024 == 0 {
		// 				c++
		// 			}
		// 		}
		// 		return 1234 * (c - 1)
		// 	}(),
		// },
		{
			"test case 3, panic probability 0, 10401650000 ns/op",
			args{
				gCount:       2048,
				hCount:       1 << 18,
				concurrently: true,
				mod:          math.MaxInt64,
			},
			func() int64 {
				c := int64(0)
				for i := 0; i != 1<<18; i++ {
					if i%math.MaxInt64 == 0 {
						c++
					}
				}
				return 2048*1<<18 - 2048*(c-1)
			}(),
			func() int32 {
				c := int32(0)
				for i := 0; i != 1<<18; i++ {
					if i%math.MaxInt64 == 0 {
						c++
					}
				}
				return 2048 * (c - 1)
			}(),
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			RecoverAtGoroutine(tt.args.gCount, tt.args.hCount, tt.args.mod, tt.args.concurrently)
		}
	}
}

func BenchmarkRecoverAtHandler(b *testing.B) {
	type args struct {
		gCount       int
		hCount       int
		mod          int
		concurrently bool
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 int32
	}{
		// TODO: Add test cases.
		// {
		// 	"test case 1, panic probability 1/7, 105206750 ns/op",
		// 	args{
		// 		gCount:       1234,
		// 		hCount:       4321,
		// 		concurrently: true,
		// 		mod:          7,
		// 	},
		// 	func() int32 {
		// 		c := int32(0)
		// 		for i := 0; i != 4321; i++ {
		// 			if i%7 == 0 {
		// 				c++
		// 			}
		// 		}
		// 		return 1234*4321 - 1234*(c-1)
		// 	}(),
		// 	func() int32 {
		// 		c := int32(0)
		// 		for i := 0; i != 4321; i++ {
		// 			if i%7 == 0 {
		// 				c++
		// 			}
		// 		}
		// 		return 1234 * (c - 1)
		// 	}(),
		// },
		// {
		// 	"test case 2, panic probability 1/1024, 106323110 ns/op",
		// 	args{
		// 		gCount:       1234,
		// 		hCount:       4321,
		// 		concurrently: true,
		// 		mod:          1024,
		// 	},
		// 	func() int32 {
		// 		c := int32(0)
		// 		for i := 0; i != 4321; i++ {
		// 			if i%1024 == 0 {
		// 				c++
		// 			}
		// 		}
		// 		return 1234*4321 - 1234*(c-1)
		// 	}(),
		// 	func() int32 {
		// 		c := int32(0)
		// 		for i := 0; i != 4321; i++ {
		// 			if i%1024 == 0 {
		// 				c++
		// 			}
		// 		}
		// 		return 1234 * (c - 1)
		// 	}(),
		// },
		{
			"test case 3, panic probability 0, 8898063600 ns/op",
			args{
				gCount:       2048,
				hCount:       1 << 18,
				concurrently: true,
				mod:          math.MaxInt64,
			},
			func() int64 {
				c := int64(0)
				for i := 0; i != 1<<18; i++ {
					if i%math.MaxInt64 == 0 {
						c++
					}
				}
				return 2048*1<<18 - 2048*(c-1)
			}(),
			func() int32 {
				c := int32(0)
				for i := 0; i != 1<<18; i++ {
					if i%math.MaxInt64 == 0 {
						c++
					}
				}
				return 2048 * (c - 1)
			}(),
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			RecoverAtHandler(tt.args.gCount, tt.args.hCount, tt.args.mod, tt.args.concurrently)
		}
	}
}
