package randfoo

import (
	"fmt"
	"math/rand"
)

// rand result is just affected by seed, but not code(re-compile or refactor)
func RandSlice(seed int64, otherInfo string) {
	rand.Seed(seed)
	for index := 0; index != 10; index++ {
		fmt.Printf("rand: %v\n", rand.Intn(100))
	}
}
