package methodfoo

import "testing"

func TestNilPointerReceiver(t *testing.T) {
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
			NilPointerReceiver()
		})
	}
}
