package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Bencher(count int, f func(int), concurrently bool) {
	var wg sync.WaitGroup
	if concurrently {
		wg = sync.WaitGroup{}
		wg.Add(count)
	}
	t1 := time.Now()
	for index := 0; index != count; index++ {
		if concurrently {
			go func() {
				f(index)
				wg.Done()
			}()
		} else {
			f(index)
		}
	}
	if concurrently {
		wg.Wait()
	}
	t2 := time.Now()
	fmt.Printf("using time: %v milli-seconds\n", t2.Sub(t1).Milliseconds())
}

func main() {
	Bencher(1, func(index int) {
		// structfoo.StructThisMemberDiff()
		// structfoo.DerivativeWithPointerBase()
		// structfoo.BaseStructTrace()
		// structfoo.SubStructAssign()
		// structfoo.SubStructDerivative()

		// typeDeclarationContent := "[][]map[Float]map[A.Int][]*B.Int"
		// d := extractorfoo.ExtractGoVariableTypeDeclaration(typeDeclarationContent)
		// d.Traversal(0)

		// algorithmfoo.ConvertCamelCase2SnakeCaseWithPhrase

		// httpfoo.RequestExample(index)
	}, false)
}
