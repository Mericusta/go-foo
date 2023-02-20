package structfoo

import (
	"fmt"
	substruct "go-foo/src/struct-foo/sub-struct"
)

func SwapStructValueOneLine() {
	s := &struct {
		A int
		B int
	}{
		A: 1,
		B: 2,
	}

	fmt.Printf("struct s.A = %v, s.B = %v\n", s.A, s.B)
	s.A, s.B = s.B, s.A
	fmt.Printf("after one-line swap struct s.A = %v, s.B = %v\n", s.A, s.B)
}

type stmd struct {
	mv map[int]int
	sv []int
	v  int
}

func (s *stmd) GetPointerThisV() (map[int]int, []int, int) {
	return s.mv, s.sv, s.v
}

func (s stmd) GetCopyThisV() (map[int]int, []int, int) {
	return s.mv, s.sv, s.v
}

func StructThisMemberDiff() {
	mv := map[int]int{1: 1}
	sv := []int{10, 10}
	v := 100100

	ps := &stmd{
		mv: mv,
		sv: sv,
		v:  v,
	}

	pMV, pSV, pV := ps.GetPointerThisV()
	pMV[1] = 2
	pSV[0] = 20
	pV += pV

	fmt.Printf("mv = %v, sv = %v, v = %v\n", mv, sv, v)
	fmt.Printf("ps.mv = %v, ps.sv = %v, ps.v = %v\n", ps.mv, ps.sv, ps.v)

	mv = map[int]int{1: 1}
	sv = []int{10, 10}
	v = 100100

	cs := &stmd{
		mv: mv,
		sv: sv,
		v:  v,
	}
	cMV, cSV, cV := cs.GetPointerThisV()
	cMV[1] = 2
	cSV[0] = 20
	cV += cV

	fmt.Printf("mv = %v, sv = %v, v = %v\n", mv, sv, v)
	fmt.Printf("cs.mv = %v, cs.sv = %v, cs.v = %v\n", cs.mv, cs.sv, cs.v)
}

type bInterface interface {
	Output() int
	Input(int)
}

type base struct {
	bV int
	bM map[int]int
}

func (b *base) Output() int {
	return b.bV
}

func (b *base) Input(i int) {
	b.bV = i
}

type derivative struct {
	base
	dV int
}

func (d *derivative) Output() int {
	return d.dV
}

func (b *derivative) Input(i int) {
	b.base.Input(i)
	b.dV = i
	b.base.bM = make(map[int]int)
	b.base.bM[i] = i
}

func (b *derivative) ModBMap(k, v int) {
	b.bM = make(map[int]int)
	b.bM[k] = v
}

type otherDerivative struct {
	base
	oV int
}

func DerivativeWithPointerBase() {
	var i1, i2, i3 bInterface
	i1 = &base{bV: 1}
	fmt.Printf("base i = %v\n", i1.Output()) // 1
	i1.Input(10)
	fmt.Printf("base i = %v\n", i1.Output()) // 10

	i2 = &derivative{base: base{bV: 1}, dV: 10}
	fmt.Printf("derivative i = %v\n", i2.Output()) // 10
	i2.Input(20)
	fmt.Printf("derivative i = %v\n", i2.Output())                         // 20
	fmt.Printf("derivative i.base = %v\n", i2.(*derivative).base.Output()) // 20
	fmt.Printf("derivative i.base = %v\n", i2.(*derivative).base.bM)       // 20
	d := &derivative{base: base{bV: 1, bM: map[int]int{1: 10}}, dV: 10}
	d.ModBMap(30, 30)
	fmt.Printf("d = %v\n", d.bM)

	i3 = &otherDerivative{base: base{bV: 2}, oV: 20}
	fmt.Printf("otherDerivative i = %v\n", i3.Output()) // 2
	i3.Input(30)
	fmt.Printf("otherDerivative i = %v\n", i3.Output()) // 30
}

func newBase() base {
	b := base{}
	fmt.Printf("newBase, %p\n", &b)
	return b
}

func newBasePointer() *base {
	b := &base{}
	fmt.Printf("newBase, %p\n", b)
	return b
}

type derivativeStruct struct {
	base
}

type derivativePointer struct {
	*base
}

func BaseStructTrace() {
	b := newBase()                 // p1
	fmt.Printf("d.base, %p\n", &b) // p2
	bp := newBasePointer()         // p3
	fmt.Printf("d.base, %p\n", bp) // p3

	d := derivativeStruct{
		base: b,
	}
	dBase := d.base
	fmt.Printf("dBase, %p\n", &dBase)

	dp := derivativePointer{
		base: bp,
	}
	dpBase := dp.base
	fmt.Printf("dpBase, %p\n", &dpBase)
}

func SubStructAssign() {
	ss := substruct.SubStruct{}
	ss.Assign(10)                           // 10
	fmt.Printf("ss.Val() = %v\n", ss.Val()) // 0
}

type subStructDerivative struct {
	substruct.SubStruct
}

func SubStructDerivative() {
	s := subStructDerivative{}
	s.SubStruct.PubV = 10
	fmt.Printf("PubV = %v\n", s.GetPubV())
}
