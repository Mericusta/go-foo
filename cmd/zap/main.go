package main

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	var err error
	loggerCfg := zap.NewDevelopmentConfig()
	loggerCfg.OutputPaths = []string{"stdout"}
	// loggerCfg.OutputPaths = append(loggerCfg.OutputPaths, "stdout")
	logger, err = loggerCfg.Build()
	if err != nil {
		panic(err)
	}
}

type s struct {
	V      int         `json:"V,omitempty"`
	IM     map[int]int `json:"IM,omitempty"`
	IS     IS          `json:"IS,omitempty"`
	SS     *s          `json:"SS,omitempty"`
	locker *sync.RWMutex
}

// MarshalLogObject implements zapcore.ObjectMarshaler.
func (s *s) MarshalLogObject(e zapcore.ObjectEncoder) error {
	s.locker.RLock()
	e.AddInt("v", s.V)
	// e.Add("IM", s.IM)
	e.AddArray("IS", s.IS)
	// e.AddInt("SS", s.SS)
	s.locker.RUnlock()
	return nil
}

type IS []int

func (is IS) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range is {
		el := newErrArrayElem(is[i])
		err := arr.AppendObject(el)
		el.Free()
		if err != nil {
			return err
		}
	}
	return nil
}

type Pool[T any] struct {
	pool sync.Pool
}

func New[T any](fn func() T) *Pool[T] {
	return &Pool[T]{
		pool: sync.Pool{
			New: func() any {
				return fn()
			},
		},
	}
}

func (p *Pool[T]) Get() T {
	return p.pool.Get().(T)
}

// Put returns x into the pool.
func (p *Pool[T]) Put(x T) {
	p.pool.Put(x)
}

var _intArrayElemPool = New(func() *intArrayElem {
	return &intArrayElem{}
})

type intArrayElem struct{ i int }

func newErrArrayElem(i int) *intArrayElem {
	e := _intArrayElemPool.Get()
	e.i = i
	return e
}

func (e *intArrayElem) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	return arr.AppendObject(e)
}

func (e *intArrayElem) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return nil
}

func (e *intArrayElem) Free() {
	e.i = 0
	_intArrayElemPool.Put(e)
}

// func (s *s) String() string {
// 	b, _ := json.Marshal(s)
// 	return string(b)
// }

func main() {
	var (
		parent *s
		child  *s
		c      chan *s
		wCount int
		rCount int
		wwg    *sync.WaitGroup
		rwg    *sync.WaitGroup
	)

	child = &s{}
	child.V = 10
	child.IM = make(map[int]int)
	child.IS = make([]int, 0, 8)
	child.locker = &sync.RWMutex{}

	parent = &s{}
	parent.V = 0
	parent.IM = make(map[int]int)
	parent.IS = make([]int, 0, 8)
	parent.SS = child
	parent.locker = &sync.RWMutex{}

	wCount = 16
	wwg = &sync.WaitGroup{}
	wwg.Add(wCount)
	rCount = 1
	rwg = &sync.WaitGroup{}
	rwg.Add(rCount)
	c = make(chan *s, wCount)

	writeG := func(i int, v *s) {
		v.locker.Lock()
		v.V = i
		v.IM[v.V] = i
		v.IS = append(v.IS, i)
		v.SS.V = i
		v.SS.IM[v.SS.V] = i * 10
		v.SS.IS = append(v.SS.IS, i*10)
		v.locker.Unlock()
		c <- v
		logger.Info("writeG", zap.Int("i", i), zap.Object("v", v))
		wwg.Done()
	}

	guardG := func() {
		wwg.Wait()
		close(c)
	}

	readG := func(i int) {
		for v := range c {
			// v.locker.RLock()
			logger.Info("readG", zap.Int("i", i), zap.Any("v", v))
			// v.locker.RUnlock()
		}
		rwg.Done()
	}

	go guardG()
	for index := 0; index != rCount; index++ {
		go readG(index)
	}
	for index := 0; index != wCount; index++ {
		go writeG(index, parent)
	}

	rwg.Wait()
}
