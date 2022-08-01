package interfacefoo

import (
	"testing"
)

func BenchmarkEmptyInterface(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			EmptyInterface()
		}
	}
}
