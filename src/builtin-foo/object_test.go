package builtinfoo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGoroutinePassObjectPointerFoo(t *testing.T) {
	type args struct {
		generatePointer bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{generatePointer: false},
		},
		{
			"test case 1",
			args{generatePointer: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GoroutinePassObjectPointerFoo(tt.args.generatePointer)
		})
	}
}

func TestConvertAnyObjectToByteArray(t *testing.T) {
	type args struct {
		o *es
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				o: &es{
					i: 1,
					s: 2,
					f: 3,
				},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertAnyObjectToByteArray(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertAnyObjectToByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertType2(t *testing.T) {
	tests := []struct {
		name  string
		want  bool
		want1 []byte
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			true,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := convertType2()
			if got != tt.want {
				t.Errorf("convertType2() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				fmt.Printf("got1 = %+v, ptr = %p\n", got1, got1)
				fmt.Printf("got1 = %+v\n", convertByteArrayToObject(got1))
				t.Errorf("convertType2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStructMemoryAlignCalculateProcess(t *testing.T) {
	type args struct {
		compilerDefaultAlign int
		smDesc               []*structMemberSizeAlignTypeDesc
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 []int
		want3 string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				8,
				[]*structMemberSizeAlignTypeDesc{
					{size: 16, align: 8, desc: "string"},
					{size: 1, align: 1, desc: "bool"},
					{size: 2, align: 2, desc: "int16"},
				},
			},
			24,
			5,
			[]int{16, 1, 3},
			"[ 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 | 0 | _ 0 0 | _ _ _ _ ]",
		},
		{
			"test case 2",
			args{
				8,
				[]*structMemberSizeAlignTypeDesc{
					{size: 1, align: 1, desc: "bool"},
					{size: 4, align: 4, desc: "int32"},
					{size: 1, align: 1, desc: "int8"},
					{size: 8, align: 8, desc: "int64"},
					{size: 1, align: 1, desc: "byte"},
				},
			},
			32,
			17,
			[]int{1, 7, 1, 15, 1},
			"[ 0 | _ _ _ 0 0 0 0 | 0 | _ _ _ _ _ _ _ 0 0 0 0 0 0 0 0 | 0 | _ _ _ _ _ _ _ ]",
		},
		{
			"test case 3",
			args{
				8,
				[]*structMemberSizeAlignTypeDesc{
					{size: 1, align: 1, desc: "bool"},
					{size: 2, align: 2, desc: "int16"},
					{size: 24, align: 8, desc: "[]byte"},
					{size: 8, align: 8, desc: "int64"},
				},
			},
			40,
			5,
			[]int{1, 3, 28, 8},
			"[ 0 | _ 0 0 | _ _ _ _ 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 | 0 0 0 0 0 0 0 0 ]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := StructMemoryAlignCalculateProcess(tt.args.compilerDefaultAlign, tt.args.smDesc)
			if got != tt.want {
				t.Errorf("StructMemoryAlignCalculateProcess() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("StructMemoryAlignCalculateProcess() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("StructMemoryAlignCalculateProcess() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("StructMemoryAlignCalculateProcess() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}

func Test_convertStruct2Example(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{c: 1 << 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convertStruct2Example(tt.args.c)
		})
	}
}

func TestConvertStringToStringStruct2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *stringStruct
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{
				s: "I am a boy,You are a girl,We are human",
			},
			want: &stringStruct{
				s1: "I am a boy",
				s2: "You are a girl",
				s3: "We are human",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertStringToStringStruct2(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertStringToStringStruct2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertStringToStringStruct1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *stringStruct
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{
				s: "I am a boy,You are a girl,We are human",
			},
			want: &stringStruct{
				s1: "I am a boy",
				s2: "You are a girl",
				s3: "We are human",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertStringToStringStruct1(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertStringToStringStruct1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertStringToStringStruct0(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *stringStruct
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{
				s: "I am a boy,You are a girl,We are human",
			},
			want: &stringStruct{
				s1: "I am a boy",
				s2: "You are a girl",
				s3: "We are human",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertStringToStringStruct0(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertStringToStringStruct0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoTypeSizeAndAlign(t *testing.T) {
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
			GoTypeSizeAndAlign()
		})
	}
}

func Test_TConvertObjectToByteArray_b3541ecaa67f47e73394172a8534703e(t *testing.T) {
	type args struct {
		o *[]byte
	}
	tests := []struct {
		name  string
		args  args
		want0 []byte
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := TConvertObjectToByteArray(tt.args.o)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("TConvertObjectToByteArray() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}
