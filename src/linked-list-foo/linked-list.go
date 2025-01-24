package linkedlistfoo

import "container/list"

// 非侵入式链表 Non-intrusive linked list

type nonIntrusiveLinkedListNode[T any] struct {
	prev  *nonIntrusiveLinkedListNode[T]
	next  *nonIntrusiveLinkedListNode[T]
	value T
}

type nonIntrusiveDoubleLinkedList[T any] struct {
	head *nonIntrusiveLinkedListNode[T]
	tail *nonIntrusiveLinkedListNode[T]
}

func (l *nonIntrusiveDoubleLinkedList[T]) append(v T) { /* bala bala bala */ }
func (l *nonIntrusiveDoubleLinkedList[T]) delete(v T) { /* bala bala bala */ }
func (l *nonIntrusiveDoubleLinkedList[T]) search(v T) { /* bala bala bala */ }

// 侵入式链表 intrusive linked list

type intrusiveLinkedListNode struct {
	prev iIntrusiveLinkedListNode
	next iIntrusiveLinkedListNode
}

func (n *intrusiveLinkedListNode) GetPrev() iIntrusiveLinkedListNode {
	return n.prev
}

func (n *intrusiveLinkedListNode) SetPrev(prev iIntrusiveLinkedListNode) {
	n.prev = prev
}

func (n *intrusiveLinkedListNode) GetNext() iIntrusiveLinkedListNode {
	return n.next
}

func (n *intrusiveLinkedListNode) SetNext(next iIntrusiveLinkedListNode) {
	n.next = next
}

type iIntrusiveLinkedListNode interface {
	GetPrev() iIntrusiveLinkedListNode
	SetPrev(iIntrusiveLinkedListNode)
	GetNext() iIntrusiveLinkedListNode
	SetNext(iIntrusiveLinkedListNode)
	Equal(iIntrusiveLinkedListNode) bool
}

type intrusiveDoubleLinkedList[T iIntrusiveLinkedListNode] struct {
	head iIntrusiveLinkedListNode
	tail iIntrusiveLinkedListNode
}

func (l *intrusiveDoubleLinkedList[T]) append(v iIntrusiveLinkedListNode) {
	if l.head == nil || l.tail == nil {
		l.head, l.tail = v, v
		return
	}

	v.SetPrev(l.tail)
	l.tail.SetNext(v)
}

func (l *intrusiveDoubleLinkedList[T]) delete(v iIntrusiveLinkedListNode) {
	if l.head == v {
		if l.tail == v {
			l.head = nil
			l.tail = nil
			return
		}

		l.head = l.head.GetNext()
		l.head.SetPrev(nil)
		v.SetPrev(nil)
		v.SetNext(nil)
		return
	}

	if l.tail == v {
		l.tail = l.tail.GetPrev()
		l.tail.SetNext(nil)
		v.SetPrev(nil)
		v.SetNext(nil)
		return
	}

	prev := v.GetPrev()
	next := v.GetNext()
	prev.SetNext(next)
	next.SetPrev(prev)
	v.SetPrev(nil)
	v.SetNext(nil)
}

// intrusive linked list no need search

type myStruct1 struct {
	*intrusiveLinkedListNode
	iv int
}

func (s *myStruct1) Equal(is iIntrusiveLinkedListNode) bool {
	_s, ok := is.(*myStruct1)
	if !ok {
		return false
	}
	return s.iv == _s.iv
}

type myStruct2 struct {
	*intrusiveLinkedListNode
	sv string
}

func (s *myStruct2) Equal(is iIntrusiveLinkedListNode) bool {
	_s, ok := is.(*myStruct2)
	if !ok {
		return false
	}
	return s.sv == _s.sv
}

type myStruct3 struct {
	*intrusiveLinkedListNode
	s1 *myStruct1
}

func (s *myStruct3) Equal(is iIntrusiveLinkedListNode) bool {
	_s, ok := is.(*myStruct3)
	if !ok {
		return false
	}
	return s.s1.Equal(_s.s1)
}

type myStruct4[T comparable] struct {
	*intrusiveLinkedListNode
	t T
}

func (s *myStruct4[T]) Equal(is iIntrusiveLinkedListNode) bool {
	_s, ok := is.(*myStruct4[T])
	if !ok {
		return false
	}
	return s.t == _s.t
}

func intrusiveDoubleLinkedListFoo() {
	l := &intrusiveDoubleLinkedList[iIntrusiveLinkedListNode]{}
	s1 := &myStruct1{
		intrusiveLinkedListNode: &intrusiveLinkedListNode{},
		iv:                      1,
	}
	s2 := &myStruct2{
		intrusiveLinkedListNode: &intrusiveLinkedListNode{},
		sv:                      "2",
	}
	s3 := &myStruct3{
		intrusiveLinkedListNode: &intrusiveLinkedListNode{},
		s1:                      s1,
	}
	s4 := &myStruct4[*myStruct3]{
		intrusiveLinkedListNode: &intrusiveLinkedListNode{},
		t:                       s3,
	}
	l.append(s1)
	l.append(s2)
	l.append(s3)
	l.append(s4)

	list.New()
}
