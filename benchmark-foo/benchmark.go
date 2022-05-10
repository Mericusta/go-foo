package benchmarkfoo

import "fmt"

func returnMapFunction(count int) map[int]int {
	m := make(map[int]int, count)
	for index := 0; index != count; index++ {
		m[index] = index
	}
	return m
}

type element struct {
	k int
	v int
}

func returnSliceFunction(count int) []*element {
	s := make([]*element, 0, count)
	for index := 0; index != count; index++ {
		s = append(s, &element{
			k: index,
			v: index,
		})
	}
	return s
}

func returnArrayFunction(count int) []*element {
	a := make([]*element, count)
	for index := 0; index != count; index++ {
		a[index] = &element{
			k: index,
			v: index,
		}
	}
	return a
}

func passMapFunction(m map[int]int) {
	for k, v := range m {
		m[k] = v * 10
	}
}

func passSliceFunction(s []*element) {
	for _, e := range s {
		e.v = e.v * 10
	}
}

func passArrayFunction(a []*element) {
	for _, e := range a {
		e.v = e.v * 10
	}
}

func mapFunction(count int) {
	m := returnMapFunction(count)
	passMapFunction(m)
}

func sliceFunction(count int) {
	s := returnSliceFunction(count)
	passSliceFunction(s)
}

func arrayFunction(count int) {
	s := returnArrayFunction(count)
	passArrayFunction(s)
}

// ----------------------------------------------------------------

func makeRewards() map[int32]int64 {
	return map[int32]int64{
		1000: 1,
		1001: 1,
	}
}

type bagItem struct {
	ItemID    int32
	ItemCount int64
	IsNew     bool
}

func makeRewardsOpt() []*bagItem {
	return []*bagItem{
		{
			ItemID:    1000,
			ItemCount: 1,
			IsNew:     false,
		},
		{
			ItemID:    1001,
			ItemCount: 1,
			IsNew:     false,
		},
	}
}

func makeRewardGroupDataWithTag() map[int32]int64 {
	rewardResult := make(map[int32]int64)
	resMap := makeRewards()
	rewardResult = mergeRewardResult(rewardResult, resMap)
	return rewardResult
}

func makeRewardGroupDataWithTagOpt() []*bagItem {
	rewardResult := make([]*bagItem, 0, 4)
	resMap := makeRewardsOpt()
	rewardResult = mergeRewardResultOpt(rewardResult, resMap)
	return rewardResult
}

func doRewardJackPotType5() map[int32]int64 {
	return makeRewardGroupDataWithTag()
}

func doRewardJackPotType5Opt() []*bagItem {
	return makeRewardGroupDataWithTagOpt()
}

func mergeRewardResult(origin, new map[int32]int64) map[int32]int64 {
	for t, r := range new {
		origin[t] += r
	}
	return origin
}

func mergeRewardResultOpt(origin, new []*bagItem) []*bagItem {
	for i := 0; i != len(new); i++ {
		has := false
		for j := 0; j != len(origin); j++ {
			if origin[j].ItemID == new[i].ItemID {
				origin[j].ItemCount += new[i].ItemCount
				has = true
				break
			}
		}
		if !has {
			origin = append(origin, new[i])
		}
	}
	return origin
}

func GetRewardItemByPoolID() map[int32]int64 {
	rewardResult := make(map[int32]int64)
	resMap := doRewardJackPotType5()
	rewardResult = mergeRewardResult(rewardResult, resMap)
	return rewardResult
}

func GetRewardItemByPoolIDOpt() []*bagItem {
	rewardResult := make([]*bagItem, 0, 4)
	resMap := doRewardJackPotType5Opt()
	rewardResult = mergeRewardResultOpt(rewardResult, resMap)
	return rewardResult
}

func Pray() map[int32]int64 {
	// prepare
	prayTimes := 10
	m := make(map[int32]int64)

	// get reward
	for prayIndex := 0; prayIndex < prayTimes; prayIndex++ {
		for k, v := range GetRewardItemByPoolID() {
			m[k] += v
		}
	}
	return m
}

func PrayOpt() []*bagItem {
	// prepare
	prayTimes := 10
	s := make([]*bagItem, 0, 10)
	// get rewrad
	for prayIndex := 0; prayIndex < prayTimes; prayIndex++ {
		for _, item := range GetRewardItemByPoolIDOpt() {
			has := false
			for index := 0; index < len(s); index++ {
				if s[index].ItemID == item.ItemID {
					s[index].ItemCount += item.ItemCount
					has = true
					break
				}
			}
			if !has {
				s = append(s, item)
			}
		}
	}
	return s
}

func LambdaCaptureFunction(f func(int) int) int {
	if f != nil {
		return f(10)
	}
	return -1
}

func LambdaCapture(testCase int) int {
	var r int = 0
	v := LambdaCaptureFunction(func(i int) int {
		for index := 0; index != i; index++ {
			r = index
		}
		return 0
	})
	fmt.Printf("testCase = %v, v = %v, r = %v\n", testCase, v, r)
	return v
}
