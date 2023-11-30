package builtinfoo

import (
	"fmt"
	"time"
)

// defer with named function can only catch type
func DeferWithNamedFunctionFoo() {
	var localVariable int
	defer deferNamedFunc(localVariable)
	localVariable = int(time.Now().Unix())%60 + 1
	fmt.Printf("localVariable %v\n", localVariable)
}

func deferNamedFunc(v interface{}) {
	switch v.(type) {
	case int:
		if v != 0 {
			panic(v)
		}
		fmt.Printf("deferAnonymousFunc catch variable %v\n", v)
	default:
		panic(v)
	}
}

// defer with anonymous function can catch variable
func DeferWithAnonymousFunctionFoo() {
	var localVariable int
	defer func() {
		deferAnonymousFunc(localVariable)
	}()
	localVariable = int(time.Now().Unix()) % 60
	fmt.Printf("localVariable %v\n", localVariable)
}

func deferAnonymousFunc(v interface{}) {
	if v == nil {
		panic(v)
	}
	fmt.Printf("deferAnonymousFunc catch variable %v\n", v)
}

func deferRecoverInFunc() {
	recoverFunc := func() {
		if p := recover(); p != nil {
			fmt.Println("recoverFunc, p =", p)
		}
	}

	defer func() {
		recoverFunc()
		if p := recover(); p != nil {
			fmt.Println("recoverFunc, p =", p)
		}
	}()

	panic("panic here")
}
