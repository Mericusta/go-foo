package goroutinefoo

import (
	"testing"
)

func TestOpenSoMuchGoRoutine(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OpenSoMuchGoRoutine()
		})
	}
}

func TestAllMIsWorking(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AllMIsWorking()
		})
	}
}

func TestGoroutinePanicAndRecoverFoo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GoroutinePanicAndRecoverFoo()
		})
	}
}

// func TestRecoverAtHandler(t *testing.T) {
// 	type args struct {
// 		gCount       int
// 		hCount       int
// 		concurrently bool
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int32
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			"test case 1",
// 			args{
// 				gCount:       1,
// 				hCount:       100,
// 				concurrently: false,
// 			},
// 			func() int32 {
// 				c := int32(0)
// 				for i := 0; i != 100; i++ {
// 					if i%7 == 0 {
// 						c++
// 					}
// 				}
// 				return 1 * (c - 1)
// 			}(),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := RecoverAtHandler(tt.args.gCount, tt.args.hCount, tt.args.concurrently); got != tt.want {
// 				t.Errorf("RecoverAtHandler() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRecoverAtGoroutine(t *testing.T) {
// 	type args struct {
// 		gCount       int
// 		hCount       int
// 		concurrently bool
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int32
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			"test case 1",
// 			args{
// 				gCount:       1,
// 				hCount:       100,
// 				concurrently: false,
// 			},
// 			func() int32 {
// 				c := int32(0)
// 				for i := 0; i != 100; i++ {
// 					if i%7 == 0 {
// 						c++
// 					}
// 				}
// 				return 1 * (c - 1)
// 			}(),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := RecoverAtGoroutine(tt.args.gCount, tt.args.hCount, tt.args.concurrently); got != tt.want {
// 				t.Errorf("RecoverAtGoroutine() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestRecoverAtHandler(t *testing.T) {
	type args struct {
		gCount       int
		hCount       int
		concurrently bool
	}
	tests := []struct {
		name  string
		args  args
		want  int32
		want1 int32
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				gCount:       1234,
				hCount:       4321,
				concurrently: true,
			},
			func() int32 {
				c := int32(0)
				for i := 0; i != 4321; i++ {
					if i%7 == 0 {
						c++
					}
				}
				return 1234*4321 - 1234*(c-1)
			}(),
			func() int32 {
				c := int32(0)
				for i := 0; i != 4321; i++ {
					if i%7 == 0 {
						c++
					}
				}
				return 1234 * (c - 1)
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RecoverAtHandler(tt.args.gCount, tt.args.hCount, tt.args.concurrently)
			if got != tt.want {
				t.Errorf("RecoverAtHandler() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("RecoverAtHandler() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRecoverAtGoroutine(t *testing.T) {
	type args struct {
		gCount       int
		hCount       int
		concurrently bool
	}
	tests := []struct {
		name  string
		args  args
		want  int32
		want1 int32
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				gCount:       1,
				hCount:       10,
				concurrently: false,
			},
			func() int32 {
				c := int32(0)
				for i := 0; i != 10; i++ {
					if i%7 == 0 {
						c++
					}
				}
				return 1*10 - 1*(c-1)
			}(),
			func() int32 {
				c := int32(0)
				for i := 0; i != 10; i++ {
					if i%7 == 0 {
						c++
					}
				}
				return 1 * (c - 1)
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RecoverAtGoroutine(tt.args.gCount, tt.args.hCount, tt.args.concurrently)
			if got != tt.want {
				t.Errorf("RecoverAtGoroutine() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("RecoverAtGoroutine() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRecoverAtGoroutineCorrectly(t *testing.T) {
	type args struct {
		gCount       int
		hCount       int
		concurrently bool
	}
	tests := []struct {
		name  string
		args  args
		want  int32
		want1 int32
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				gCount:       1234,
				hCount:       4321,
				concurrently: true,
			},
			func() int32 {
				c := int32(0)
				for i := 0; i != 4321; i++ {
					if i%7 == 0 {
						c++
					}
				}
				return 1234*4321 - 1234*(c-1)
			}(),
			func() int32 {
				c := int32(0)
				for i := 0; i != 4321; i++ {
					if i%7 == 0 {
						c++
					}
				}
				return 1234 * (c - 1)
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RecoverAtGoroutineCorrectly(tt.args.gCount, tt.args.hCount, tt.args.concurrently)
			if got != tt.want {
				t.Errorf("RecoverAtGoroutineCorrectly() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("RecoverAtGoroutineCorrectly() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
