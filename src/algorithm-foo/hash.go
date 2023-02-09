package algorithmfoo

import (
	"fmt"

	optimus "github.com/pjebs/optimus-go"
)

// ----------------------------------------------------------------

// Knuth's Hashing Algorithm Implementation: optimus-go

const PrimeNum = 233323327
const RandNum = 214748365

var optimusPrime optimus.Optimus

func EncodeID(ID int64) int64 {
	optimusPrime = optimus.NewCalculated(PrimeNum, RandNum)
	return int64(optimusPrime.Encode(uint64(ID)))
}

func DecodeID(identifier int64) int64 {
	optimusPrime = optimus.NewCalculated(PrimeNum, RandNum)
	return int64(optimusPrime.Decode(uint64(identifier)))
}

// ----------------------------------------------------------------

func DynamicHashAverageAlgorithm(nCount, sCount, rCount, luckyS int) {
	numSlice := make([]int, 0, nCount)
	for i := 1; i <= nCount; i++ {
		numSlice = append(numSlice, i)
	}
	slotHashResult := make(map[int]int)
	for i := 1; i <= sCount; i++ {
		slotHashResult[i] = 0
	}

	for i := 0; i < rCount; i++ {
		toHashNumSlice := PutAwayRandom(nCount/2, func(f func(int, int) bool) {
			for _, n := range numSlice {
				if n%sCount+1 == luckyS {
					if !f(n, 2) {
						return
					}
				} else {
					if !f(n, 1) {
						return
					}
				}
			}
		})
		for _, n := range toHashNumSlice {
			slotHashResult[modHash(n, sCount)+1]++
		}
	}

	fmt.Printf("slot hash result %v\n", slotHashResult)
}

func modHash(n, s int) int {
	return n % s
}
