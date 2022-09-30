package main

import (
	"fmt"
)

func main() {
	var fs [4]func()
	{
	}

	var v *int

	for i := 0; i < 4; i++ {
		fs[i] = func() {
			fmt.Println("打印v = ", i)
			i *= 10
		}
		v = &i
		fmt.Printf("i addr %v\n", &i)
	}

	for _, f := range fs {
		f()
	}
	fmt.Printf("v addr = %v, v value = %v\n", v, *v)
}
