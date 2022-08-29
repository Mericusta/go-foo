package methodfoo

import "fmt"

type NilePointerStruct struct{ v int }

func (s *NilePointerStruct) nilPointerReceiverFunc() {
	fmt.Printf("on nilPointerReceiverFunc\n")
	if s != nil {
		panic("not nil")
	}
	fmt.Printf("try to use nil pointer receiver %v will panic\n", s.v)
}

func NilPointerReceiver() {
	var ns *NilePointerStruct
	ns.nilPointerReceiverFunc() // output 'on nilPointerReceiverFunc' then panic
}
