package builtinfoo

import (
	"testing"
)

func BenchmarkDeferWithAnonymousFunctionFoo(b *testing.B) {
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
			DeferWithAnonymousFunctionFoo()
		}
	}
}

func BenchmarkDeferWithNamedFunctionFoo(b *testing.B) {
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
			DeferWithNamedFunctionFoo()
		}
	}
}
