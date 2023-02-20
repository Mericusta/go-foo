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
	(*(*T)(unsafe.Pointer(a))).Eat()
}

type Lion struct {
	animal[Lion]
}

func (l Lion) Eat() {
	fmt.Printf("lion eat")
}

type Cat struct {
	animal[Cat]
}

func (l Cat) Eat() {
	fmt.Printf("cat eat")
}
