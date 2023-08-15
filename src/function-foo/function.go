package functionfoo

import "fmt"

type ExampleStruct struct {
	A int
	B int
}

func (e *ExampleStruct) ReturnExampleStruct() (ExampleStruct, *ExampleStruct) {
	e.A = 10
	e.B = 20
	return *e, e
}

func ReturnExampleStructTest() {
	e := &ExampleStruct{A: 100}
	fmt.Printf("e = %p\n", e)
	fmt.Printf("e.A = %p\n", &e.A)
	fmt.Printf("e.B = %p\n", &e.B)

	ce, pe := e.ReturnExampleStruct()
	fmt.Printf("ce = %p\n", &ce)
	fmt.Printf("ce.A = %p\n", &ce.A)
	fmt.Printf("ce.B = %p\n", &ce.B)

	fmt.Printf("pe = %p\n", pe)
	fmt.Printf("pe.A = %p\n", &pe.A)
	fmt.Printf("pe.B = %p\n", &pe.B)
}

func passStructFoo() {
	s := make([]struct {
		// a string
		A string
		B string
	}, 0, 8)
	for i := 0; i != 8; i++ {
		s = append(s, struct {
			// a string
			A string
			B string
		}{
			// a: fmt.Sprintf("a%v", i),
			A: fmt.Sprintf("A%v", i),
			B: fmt.Sprintf("B%v", i),
		})
	}
	passStructSlice(s)
}

// 带有未导出变量的匿名结构体数组，无法通过函数传递
func passStructSlice(s []struct {
	// a string
	A string
	B string
}) {
	for _, _s := range s {
		fmt.Printf("%v, %v\n", _s.A, _s.B)
	}
}
