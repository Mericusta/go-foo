package interfacegoo

type I interface {
	IFunc()
}

type S struct{}

func (s S) IFunc() {

}

type SP struct{}

func (s *SP) IFunc() {

}

func StructAndPointerImplement() {
	var i I

	i = S{}
	i.IFunc()

	i = &SP{}
	i.IFunc()
}
