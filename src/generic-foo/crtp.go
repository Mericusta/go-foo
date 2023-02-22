package genericfoo

import (
	"fmt"
	"unsafe"
)

type Animal interface {
	Eat()
}

type animal[T Animal] struct{}

func (a *animal[T]) Eat() {
	(*(*T)(unsafe.Pointer(a))).Eat() // need type assert in compile time
}

type Lion struct {
	animal[Lion]
}

func (l Lion) Eat() {
	fmt.Printf("lion eat\n")
}

type Cat struct {
	animal[Cat]
}

func (l Cat) Eat() {
	fmt.Printf("cat eat\n")
}

func CRTPCall() {
	var a Animal
	a = &animal[Lion]{}
	a.Eat()

	a = &animal[Cat]{}
	a.Eat()
}
