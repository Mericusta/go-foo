package structfoo

import "fmt"

func SwapStructValueOneLine() {
	s := &struct {
		A int
		B int
	}{
		A: 1,
		B: 2,
	}

	fmt.Printf("struct s.A = %v, s.B = %v\n", s.A, s.B)
	s.A, s.B = s.B, s.A
	fmt.Printf("after one-line swap struct s.A = %v, s.B = %v\n", s.A, s.B)
}
