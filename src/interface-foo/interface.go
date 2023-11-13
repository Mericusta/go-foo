package interfacefoo

import "fmt"

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

func GetEmptyInterface() interface{} {
	return nil
}

func GetIButReturnNil() I {
	return nil
}

func GetI() I {
	return &S{}
}

func GetSButReturnNil() *S {
	return nil
}

func GetS() *S {
	return &S{}
}

func EmptyInterface() {
	var i I
	fmt.Println(i == nil) // true

	i = GetIButReturnNil()
	fmt.Println(i == nil) // true

	i = GetSButReturnNil()
	fmt.Println(i == nil) // false

	i = GetI()
	fmt.Println(i == nil) // false

	i = GetS()
	fmt.Println(i == nil) // false

	var s *S
	fmt.Println(s == nil) // true

	s = GetSButReturnNil()
	fmt.Println(s == nil) // true

	s = GetS()
	fmt.Println(s == nil) // false

	nowI := GetIButReturnNil()
	fmt.Println(nowI == nil) // true

	nowS := GetSButReturnNil()
	fmt.Println(nowS == nil) // true

	var nilI I
	var nilS *S
	fmt.Println(nilI == nilS) // false
}

type InterfaceTypeAssertStruct struct{ i int }

func InterfaceTypeAssert(i interface{}) (int, bool) {
	is, ok := i.(*InterfaceTypeAssertStruct)
	if !ok {
		return 0, ok
	}
	return is.i, ok
}

type MyError1 struct {
	error
}

func newMyError1(e error) *MyError1 {
	return &MyError1{error: e}
}

type MyError2 struct {
	error
}

func newMyError2(e error) *MyError2 {
	return &MyError2{error: e}
}

func ErrorPassNilCheckFoo() {
	myError1 := newMyError1(nil)
	if myError1.error != nil {
		fmt.Printf("myError1 is not nil\n")
	}
	myError2 := newMyError2(myError1.error)
	if myError2.error != nil {
		fmt.Printf("myError2 is not nil\n")
	}
}
