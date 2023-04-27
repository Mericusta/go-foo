package mockfoo

import "fmt"

type ExampleInterface interface {
	ExampleMethod(int, string, interface{}) error
}

type UsageStruct struct {
	i ExampleInterface
}

func (s *UsageStruct) Use(p1 int, p2 string, p3 interface{}) error {
	s.i.ExampleMethod(p1, p2, p3)
	if p1%2 == 0 {
		return fmt.Errorf("not odd")
	}
	return nil
}
