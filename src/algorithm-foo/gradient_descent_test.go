package algorithmfoo_test

import (
	algorithmfoo "go-foo/src/algorithm-foo"
	"testing"
)

func TestGradientDescentFoo(t *testing.T) {
	tests := []struct {
		name string // description of this test case
	}{
		// TODO: Add test cases.
		{name: "TestGradientDescentFoo Case 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			algorithmfoo.GradientDescentFoo()
		})
	}
}
