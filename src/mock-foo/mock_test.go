package mockfoo

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestExampleMethod(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockExampleInterface := NewMockExampleInterface(mockCtrl)
	testUsageStruct := &UsageStruct{i: mockExampleInterface}

	mockExampleInterface.EXPECT().ExampleMethod(1, gomock.Any(), gomock.Any()).Return(nil).Times(1)
	mockExampleInterface.EXPECT().ExampleMethod(2, gomock.Any(), gomock.Any()).Return(fmt.Errorf("not odd")).Times(1)

	testUsageStruct.Use(1, "2", []int{3})
	testUsageStruct.Use(2, "2", []int{3})
}

func TestMockFoo(t *testing.T) {
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
			MockFoo()
		})
	}
}
