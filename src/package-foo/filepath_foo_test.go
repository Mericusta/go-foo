package packagefoo

import "testing"

func TestFileInfoForDirLink(t *testing.T) {
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
			FileInfoForDirLink()
		})
	}
}
