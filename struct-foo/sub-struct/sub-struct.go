package substruct

type SubStruct struct {
	val  int
	PubV int
}

func (s *SubStruct) Assign(v int) {
	s.val = v
}

func (s SubStruct) Val() int {
	return s.val
}

func (s SubStruct) GetPubV() int {
	return s.PubV
}
