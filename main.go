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
		// 	regexpfoo.GO_IMPORT_SCOPE_CONTENT,
		// 	regexpfoo.GO_IMPORT_SCOPE_EXPRESSION,
		// )

		// regexpfoo.RegexpTest(-1,
		// 	regexpfoo.GO_EACH_IMPORT_CONTENT,
		// 	regexpfoo.GO_EACH_IMPORT_EXPRESSION,
		// 	regexpfoo.GO_EACH_IMPORT_SUBMATCH_ALIAS,
		// 	regexpfoo.GO_EACH_IMPORT_SUBMATCH_PATH,
		// )

		// regexpfoo.RegexpTest(-1,
		// 	regexpfoo.GO_PACKAGE_SCOPE_CONTENT,
		// 	regexpfoo.GO_PACKAGE_SCOPE_EXPRESSION,
		// 	regexpfoo.GO_PACKAGE_SCOPE_EXPRESSION_SUBMATCH_NAME,
		// )

		// regexpfoo.RegexpTest(-1,
		// 	regexpfoo.GO_VARIABLE_DECLARATION_CONTENT,
		// 	regexpfoo.GO_VARIABLE_DECLARATION_EXPRESSION,
		// 	regexpfoo.GO_VARIABLE_DECLARATION_EXPRESSION_SUBMATCH_NAME,
		// 	regexpfoo.GO_VARIABLE_DECLARATION_EXPRESSION_SUBMATCH_TYPE,
		// )

		// regexpfoo.RegexpTest(-1,
		// 	regexpfoo.GO_VARIABLE_TYPE_MAP_DECLARATION_CONTENT,
		// 	regexpfoo.GO_VARIABLE_TYPE_MAP_DECLARATION_EXPRESSION,
		// 	regexpfoo.GO_VARIABLE_TYPE_MAP_DECLARATION_EXPRESSION_SUBMATCH_KEY,
		// 	regexpfoo.GO_VARIABLE_TYPE_MAP_DECLARATION_EXPRESSION_SUBMATCH_VALUE,
		// )

		regexpfoo.RegexpTest(-1,
			regexpfoo.GO_VARIABLE_TYPE_SLICE_DECLARATION_CONTENT,
			regexpfoo.GO_VARIABLE_TYPE_SLICE_DECLARATION_EXPRESSION,
			regexpfoo.GO_VARIABLE_TYPE_SLICE_DECLARATION_EXPRESSION_SUBMATCH_VALUE,
		)

		// regexpfoo.RegexpTest(-1,
		// 	regexpfoo.GO_VARIABLE_SHORT_IDENTIFIER_CONTENT,
		// 	regexpfoo.GO_VARIABLE_SHORT_IDENTIFIER_EXPRESSION,
		// 	regexpfoo.GO_VARIABLE_SHORT_IDENTIFIER_EXPRESSION_SUBMATCH_NAME,
		// 	regexpfoo.GO_VARIABLE_SHORT_IDENTIFIER_EXPRESSION_SUBMATCH_TYPE,
		// )

		// regexpfoo.RegexpTest(-1,
		// 	regexpfoo.GO_VARIABLE_TYPE_CONSTRUCTION_CONTENT,
		// 	regexpfoo.GO_VARIABLE_TYPE_CONSTRUCTION_EXPRESSION,
		// 	regexpfoo.GO_VARIABLE_TYPE_CONSTRUCTION_EXPRESSION_SUBMATCH_CALL,
		// 	regexpfoo.GO_VARIABLE_TYPE_CONSTRUCTION_EXPRESSION_SUBMATCH_FROM,
		// 	regexpfoo.GO_VARIABLE_TYPE_CONSTRUCTION_EXPRESSION_SUBMATCH_FUNC,
		// )

		// regexpfoo.RegexpTest(-1,
		// 	regexpfoo.GO_VARIABLE_TYPE_MAP_IDENTIFIER_CONTENT,
		// 	regexpfoo.GO_VARIABLE_TYPE_MAP_IDENTIFIER_EXPRESSION,
		// 	regexpfoo.GO_VARIABLE_TYPE_MAP_IDENTIFIER_SUBMATCH_KEY,
		// 	regexpfoo.GO_VARIABLE_TYPE_MAP_IDENTIFIER_SUBMATCH_VALUE,
		// )

		// regexpfoo.RegexpTest(-1,
		// 	regexpfoo.GO_VARIABLE_TYPE_SLICE_IDENTIFIER_CONTENT,
		// 	regexpfoo.GO_VARIABLE_TYPE_SLICE_IDENTIFIER_EXPRESSION,
		// 	regexpfoo.GO_VARIABLE_TYPE_SLICE_IDENTIFIER_EXPRESSION_SUBMATCH_VALUE,
		// )
	}, false)
}
