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
