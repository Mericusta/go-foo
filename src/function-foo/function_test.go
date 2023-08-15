package functionfoo

import (
	"reflect"
	"testing"
)

func TestExampleStruct_ReturnExampleStruct(t *testing.T) {
	tests := []struct {
		name  string
		e     *ExampleStruct
		want  ExampleStruct
		want1 *ExampleStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.e.ReturnExampleStruct()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExampleStruct.ReturnExampleStruct() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ExampleStruct.ReturnExampleStruct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReturnExampleStructTest(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReturnExampleStructTest()
		})
	}
}

func Test_passStructFoo(t *testing.T) {
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
			passStructFoo()
		})
	}
}
