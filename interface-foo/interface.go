package interfacefoo

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

// pass struct / this as interface

// 指针对象还是结构对象，谁符合接口不取决于接口本身或其他的约定
// 只取决于实现接口的方法的，是指针对象还是结构对象

func PassInterface(i I) {

}

func (s S) passStruct() {
	PassInterface(s)
}

func PassStruct() {
	s := S{}
	PassInterface(s)

	// this is wrong
	// sp := SP{}
	// sp must be pointer because struct SP implement interface I function IFunc by *SP
	sp := &SP{}
	PassInterface(sp)
}
