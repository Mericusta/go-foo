package randfoo

import (
	"fmt"
	"math/rand"
)

// rand result is just affected by seed, but not code(re-compile or refactor)
func RandSlice(seed int64, otherInfo string) {
	random := rand.New(rand.NewSource(seed))
	for index := 0; index != 10; index++ {
		fmt.Printf("rand: %v\n", random.Intn(100))
	}
}

func GetRandSlice(seed int64) []int {
	random := rand.New(rand.NewSource(seed))
	s := make([]int, 0, 4)
	for index := 0; index != 4; index++ {
		s = append(s, random.Intn(2))
	}
	return s
}
