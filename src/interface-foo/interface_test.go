package interfacefoo

import (
	"testing"
)

func TestEmptyInterface(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EmptyInterface()
		})
	}
}

func TestInterfaceTypeAssert(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		// TODO: Add test cases.
		{
			name:  "test case 1",
			args:  args{i: nil},
			want:  0,
			want1: false,
		},
		{
			name:  "test case 2",
			args:  args{i: InterfaceTypeAssertStruct{i: 1}},
			want:  0,
			want1: false,
		},
		{
			name:  "test case 3",
			args:  args{i: &InterfaceTypeAssertStruct{i: 1}},
			want:  1,
			want1: true,
		},
		{
			name:  "test case 4",
			args:  args{i: struct{}{}},
			want:  0,
			want1: false,
		},
		{
			name:  "test case 5",
			args:  args{i: struct{ i int }{i: 1}},
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := InterfaceTypeAssert(tt.args.i)
			if got != tt.want {
				t.Errorf("InterfaceTypeAssert() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("InterfaceTypeAssert() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
