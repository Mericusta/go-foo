package unicodefoo

import (
	"testing"
)

func BenchmarkUnicodeLengthFoo(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			UnicodeLengthFoo()
		}
	}
}
