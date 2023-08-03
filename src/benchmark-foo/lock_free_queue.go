package benchmarkfoo

import (
	"sync/atomic"
	"unsafe"
)

// Queue implements lock-free FIFO freelist based queue.
// ref: https://dl.acm.org/citation.cfm?doid=248052.248106
type Queue[T any] struct {
	head unsafe.Pointer
	tail unsafe.Pointer
	len  uint64
}

// NewQueue creates a new lock-free queue.
func NewQueue[T any]() *Queue[T] {
	head := directItem[T]{next: nil, v: getZero[T]()} // allocate a free item
	return &Queue[T]{
		tail: unsafe.Pointer(&head), // both head and tail points
		head: unsafe.Pointer(&head), // to the free item
	}
}

// Enqueue puts the given value v at the tail of the queue.
func (q *Queue[T]) Enqueue(v T) {
	i := &directItem[T]{next: nil, v: v} // allocate new item
	var last, lastnext *directItem[T]
	for {
		last = loaditem[T](&q.tail)
		lastnext = loaditem[T](&last.next)
		if loaditem[T](&q.tail) == last { // are tail and next consistent?
			if lastnext == nil { // was tail pointing to the last node?
				if casitem(&last.next, lastnext, i) { // try to link item at the end of linked list
					casitem(&q.tail, last, i) // enqueue is done. try swing tail to the inserted node
					atomic.AddUint64(&q.len, 1)
					return
				}
			} else { // tail was not pointing to the last node
				casitem(&q.tail, last, lastnext) // try swing tail to the next node
			}
		}
	}
}

// Dequeue removes and returns the value at the head of the queue.
// It returns nil if the queue is empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	var first, last, firstnext *directItem[T]
	for {
		first = loaditem[T](&q.head)
		last = loaditem[T](&q.tail)
		firstnext = loaditem[T](&first.next)
		if first == loaditem[T](&q.head) { // are head, tail and next consistent?
			if first == last { // is queue empty?
				if firstnext == nil { // queue is empty, couldn't dequeue
					return getZero[T](), false
				}
				casitem(&q.tail, last, firstnext) // tail is falling behind, try to advance it
			} else { // read value before cas, otherwise another dequeue might free the next node
				v := firstnext.v
				if casitem(&q.head, first, firstnext) { // try to swing head to the next node
					atomic.AddUint64(&q.len, ^uint64(0))
					return v, true // queue was not empty and dequeue finished.
				}
			}
		}
	}
}

// Length returns the length of the queue.
func (q *Queue[T]) Length() uint64 {
	return atomic.LoadUint64(&q.len)
}

type directItem[T any] struct {
	next unsafe.Pointer
	v    T
}

func loaditem[T any](p *unsafe.Pointer) *directItem[T] {
	return (*directItem[T])(atomic.LoadPointer(p))
}
func casitem[T any](p *unsafe.Pointer, old, new *directItem[T]) bool {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
func getZero[T any]() T {
	var result T
	return result
}
