package benchmarkfoo

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
