package unittestfoo

type AnExampleStruct struct {
	v int
}

func NewPointerFunc() *AnExampleStruct {
	return &AnExampleStruct{}
}

func NewPointerMapFunc() map[int]*AnExampleStruct {
	// return map[int]*AnExampleStruct{1: NewPointerFunc()}
	return map[int]*AnExampleStruct{1: {v: 1}}
}
