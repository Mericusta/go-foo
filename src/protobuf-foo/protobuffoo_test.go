package protobuffoo

import "testing"

// func TestMarshalStruct(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			"test case 1",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			UnmarshalNilPointer()
// 		})
// 	}
// }

// func TestInterfaceMarshalFoo(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			"test case 1",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			InterfaceMarshalFoo()
// 		})
// 	}
// }

// func TestMarshalEmptyStructFoo(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		want  int
// 		want1 int
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			"test case 1",
// 			0,
// 			12,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1 := MarshalEmptyStructFoo()
// 			if got != tt.want {
// 				t.Errorf("MarshalEmptyStructFoo() got = %v, want %v", got, tt.want)
// 			}
// 			if got1 != tt.want1 {
// 				t.Errorf("MarshalEmptyStructFoo() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }

func TestUnmarshalUnknownStructFoo(t *testing.T) {
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
			UnmarshalUnknownStructFoo()
		})
	}
}
