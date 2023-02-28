package algorithmfoo

func FullArrangementBackTracingMethod(c, s int) [][]int {
	results := make([][]int, 0)
	fullArrangementBackTracingMethodConcurrency(nil, c, c)
	return results
}

func fullArrangementBackTracingMethodConcurrency(s []int, c, r int) {
	if r == 0 {
		// fmt.Printf("s = %v\n", s)
		return
	}

	for i := 0; i < c; i++ {
		for _, v := range s {
			if i == v {
				goto NEXT
			}
		}
		go fullArrangementBackTracingMethodConcurrency(append([]int{i}, s...), c, r-1)
	NEXT:
	}
}

func FullArrangementBackTracingMethodDeepFirstSearch(c, s int) [][]int {
	results := make([][]int, 0)
	jud := make([]bool, c)
	fullArrangementBackTracingMethodDeepFirstSearch(jud, c, []int{}, 0)
	return results
}

func fullArrangementBackTracingMethodDeepFirstSearch(jud []bool, nums int, team []int, index int) {
	if index == nums {
		// fmt.Printf("s = %v\n", team)
		return
	}
	for i := 0; i < nums; i++ {
		if jud[i] {
			continue
		}
		team = append(team, i)
		jud[i] = true
		fullArrangementBackTracingMethodDeepFirstSearch(jud, nums, team, index+1)
		jud[i] = false
		team = team[:len(team)-1]
	}
}

func FullArrangementBitCompareMethod(c, s int) {

}
