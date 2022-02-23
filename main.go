package main

import (
	"fmt"
	regexpfoo "go-foo/regexp-foo"
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
		// regexpfoo.RegexpTest(-1,
		// 	regexpfoo.GO_STRUCT_MEMBER_IDENTIFIER_CONTENT,
		// 	regexpfoo.GO_STRUCT_MEMBER_IDENTIFIER_EXPRESSION,
		// 	regexpfoo.GO_STRUCT_MEMBER_SUBMATCH_NAME,
		// 	regexpfoo.GO_STRUCT_MEMBER_SUBMATCH_TYPE,
		// 	regexpfoo.GO_STRUCT_MEMBER_SUBMATCH_TYPE_FROM,
		// 	regexpfoo.GO_STRUCT_MEMBER_SUBMATCH_TYPE_META,
		// )

		// regexpfoo.RegexpTest(-1,
		// 	regexpfoo.GO_IMPORT_SCOPE_CONTENT,
		// 	regexpfoo.GO_IMPORT_SCOPE_EXPRESSION,
		// )

		// regexpfoo.RegexpTest(-1,
		// 	regexpfoo.GO_EACH_IMPORT_CONTENT,
		// 	regexpfoo.GO_EACH_IMPORT_EXPRESSION,
		// 	regexpfoo.GO_EACH_IMPORT_SUBMATCH_ALIAS,
		// 	regexpfoo.GO_EACH_IMPORT_SUBMATCH_PATH,
		// )

		regexpfoo.RegexpTest(-1,
			regexpfoo.GO_PACKAGE_SCOPE_CONTENT,
			regexpfoo.GO_PACKAGE_SCOPE_EXPRESSION,
			regexpfoo.GO_PACKAGE_SCOPE_SUBMATCH_NAME,
		)
	}, false)
}
