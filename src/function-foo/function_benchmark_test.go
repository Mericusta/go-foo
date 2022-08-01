package functionfoo

import (
	"testing"
)

func BenchmarkExampleStruct_ReturnExampleStruct(b *testing.B) {
	tests := []struct {
		name  string
		e     *ExampleStruct
		want  ExampleStruct
		want1 *ExampleStruct
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			tt.e.ReturnExampleStruct()
		}
	}
}

func BenchmarkReturnExampleStructTest(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			ReturnExampleStructTest()
		}
	}
}
