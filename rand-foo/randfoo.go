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

func GetRandSlice(seed int64) []int {
	rand.Seed(seed)
	s := make([]int, 0, 4)
	for index := 0; index != 4; index++ {
		s = append(s, rand.Intn(2))
	}
	return s
}
