package builtinfoo

func SwitchCaseValueFallthrough(c int) []int {
	s := make([]int, 0, 10)
	switch c {
	case 0:
		s = append(s, 0)
		fallthrough
	case 1:
		s = append(s, 1)
		fallthrough
	case 2:
		s = append(s, 2)
		fallthrough
	case 3:
		s = append(s, 3)
		fallthrough
	case 4:
		s = append(s, 4)
		fallthrough
	case 5:
		s = append(s, 5)
		fallthrough
	case 6:
		s = append(s, 6)
		fallthrough
	case 7:
		s = append(s, 7)
		fallthrough
	case 8:
		s = append(s, 8)
		fallthrough
	case 9:
		s = append(s, 9)
		fallthrough
	default:
	}
	return s
}

func SwitchCaseExpressionFallthrough(c int) []int {
	s := make([]int, 0, 10)
	switch {
	case 0 < c && c <= 10:
		s = append(s, 0)
		fallthrough
	case 10 < c && c <= 20:
		s = append(s, 1)
		fallthrough
	case 20 < c && c <= 30:
		s = append(s, 2)
		fallthrough
	case 30 < c && c <= 40:
		s = append(s, 3)
		fallthrough
	case 40 < c && c <= 50:
		s = append(s, 4)
		fallthrough
	case 50 < c && c <= 60:
		s = append(s, 5)
		fallthrough
	case 60 < c && c <= 70:
		s = append(s, 6)
		fallthrough
	case 70 < c && c <= 80:
		s = append(s, 7)
		fallthrough
	case 80 < c && c <= 90:
		s = append(s, 8)
		fallthrough
	case 90 < c && c <= 100:
		s = append(s, 9)
		fallthrough
	default:
	}
	return s
}
