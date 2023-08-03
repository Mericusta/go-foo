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

func Call(animal Animal) {
	animal.Eat()
}

func AInsert(a ...Animal) {

}

func CRTPCall() {
	a := &animal[Lion]{}
	a.Eat()
	Call(a)

	b := &animal[Cat]{}
	b.Eat()
	Call(b)

	AInsert(a, b)
}
