package contextfoo

import (
	"context"
	"reflect"
	"testing"
)

func Test_stopGoroutineWay1(t *testing.T) {
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
			stopGoroutineWay1()
		})
	}
}

func Test_stopGoroutineWay2(t *testing.T) {
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
			stopGoroutineWay2()
		})
	}
}

func Test_monitorMultiGoroutineWithContext(t *testing.T) {
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
			monitorMultiGoroutineWithContext()
		})
	}
}

func Test_monitorGoroutineWithContextAndValue(t *testing.T) {
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
			monitorGoroutineWithContextAndValue()
		})
	}
}

func Test_contextTreeCloseWay1(t *testing.T) {
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
			contextTreeCloseWay1()
		})
	}
}

func Test_contextTreeCloseWay2(t *testing.T) {
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
			contextTreeCloseWay2()
		})
	}
}

func TestTimeoutContextFoo(t *testing.T) {
	type args struct {
		timeoutSeconds  int
		businessSeconds int
		businessPanic   bool
	}
	tests := []struct {
		name string
		args args
		want *fooError
	}{
		// TODO: Add test cases.
		{
			"test case 1, overtime",
			args{
				timeoutSeconds:  2,
				businessSeconds: 10,
				businessPanic:   false,
			},
			&fooError{
				e: context.DeadlineExceeded.Error(),
			},
		},
		{
			"test case 2, panic",
			args{
				timeoutSeconds:  2,
				businessSeconds: 2,
				businessPanic:   true,
			},
			&fooError{
				e: context.Canceled.Error(),
			},
		},
		{
			"test case 3, complete",
			args{
				timeoutSeconds:  2,
				businessSeconds: 1,
				businessPanic:   false,
			},
			&fooError{
				e: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeoutContextFoo(tt.args.timeoutSeconds, tt.args.businessSeconds, tt.args.businessPanic); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeoutContextFoo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubContextCancelFoo(t *testing.T) {
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
			SubContextCancelFoo()
		})
	}
}

func Test_contextTreeControl(t *testing.T) {
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
			contextTreeControl()
		})
	}
}
