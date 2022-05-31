package benchmarkfoo

import (
	"reflect"
	"testing"
)

func TestPray(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Pray()
		})
	}
}

func TestLambdaCapture(t *testing.T) {
	type args struct {
		testCase int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "lambda capture",
			args: args{
				testCase: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LambdaCapture(tt.args.testCase); got != tt.want {
				t.Errorf("LambdaCapture() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElementsInSlice(t *testing.T) {
	getEvenSlice := func(c int) []int {
		s := make([]int, 0, c/2)
		for index := 0; index < c; index++ {
			if index%2 == 0 {
				s = append(s, index)
			}
		}
		return s
	}
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"8 elements slice test",
			args{
				count: 8,
			},
			getEvenSlice(8),
		},
		{
			"16 elements slice test",
			args{
				count: 16,
			},
			getEvenSlice(16),
		},
		{
			"32 elements slice test",
			args{
				count: 32,
			},
			getEvenSlice(32),
		},
		{
			"64 elements slice test",
			args{
				count: 64,
			},
			getEvenSlice(64),
		},
		{
			"128 elements slice test",
			args{
				count: 128,
			},
			getEvenSlice(128),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ElementsInSlice(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementsInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElementsInMap(t *testing.T) {
	getEvenMap := func(c int) map[int]struct{} {
		m := make(map[int]struct{}, c/2)
		for index := 0; index < c; index++ {
			if index%2 == 0 {
				m[index] = struct{}{}
			}
		}
		return m
	}
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want map[int]struct{}
	}{
		// TODO: Add test cases.
		{
			"8 elements map test",
			args{
				count: 8,
			},
			getEvenMap(8),
		},
		{
			"16 elements map test",
			args{
				count: 16,
			},
			getEvenMap(16),
		},
		{
			"32 elements map test",
			args{
				count: 32,
			},
			getEvenMap(32),
		},
		{
			"64 elements map test",
			args{
				count: 64,
			},
			getEvenMap(64),
		},
		{
			"128 elements map test",
			args{
				count: 128,
			},
			getEvenMap(128),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ElementsInMap(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementsInMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassStructFuncTest(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PassStructFuncTest()
		})
	}
}

func TestPassInterfaceFuncTest(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PassInterfaceFuncTest()
		})
	}
}

func TestFMTPrintfBenchmark(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FMTPrintfBenchmark()
		})
	}
}
