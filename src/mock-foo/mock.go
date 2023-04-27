package mockfoo

import (
	"fmt"
	"math"
	"math/rand"
)

type ExampleInterface interface {
	ExampleMethod(int, string, interface{}) error
}

type UsageStruct struct {
	i ExampleInterface
}

func (s *UsageStruct) Use(p1 int, p2 string, p3 interface{}) error {
	s.i.ExampleMethod(p1, p2, p3)
	if p1%2 == 0 {
		return fmt.Errorf("not odd")
	}
	return nil
}

// ----------------------------------------------------------------

type Mocker struct {
	// IntMocker
}

var DefaultMocker = &Mocker{}

func (m *Mocker) mock(v interface{}) interface{} {
	switch v.(type) {
	case bool:
		return rand.Intn(2) == 0
	case int:
		return int(math.MaxInt)
	case int8:
		return int8(math.MaxInt8)
	case int16:
		return int16(math.MaxInt16)
	case int32:
		return int32(math.MaxInt32)
	case int64:
		return int64(math.MaxInt64)
	case uint:
		return uint(math.MaxUint)
	case uint8:
		return uint8(math.MaxUint8)
	case uint16:
		return uint16(math.MaxUint16)
	case uint32:
		return uint32(math.MaxUint32)
	case uint64:
		return uint64(math.MaxUint64)
	case float32:
		return float32(math.MaxFloat32)
	case float64:
		return float64(math.MaxFloat64)
	case error:
		return fmt.Errorf("error")
	case string:
		return string("string")
	default:
		if s, ev := traitSlice[any](v); s != nil {
			sv := DefaultMocker.mock(ev)
			// return mockSlice[sv]()
		}
	}
	return nil
}

func traitSlice[T any](v interface{}) ([]T, T) {
	var ev T
	s, ok := v.([]T)
	if !ok {
		return nil, ev
	}
	return s, ev
}

func mockSlice[T any]() []T {
	l := 8
	s := make([]T, 0, l)
	for i := 0; i < l; i++ {
		s = append(s, mockSingle[T]())
	}
	return s
}

func mockSingle[T any]() T {
	var v T
	return DefaultMocker.mock(v).(T)
}

func mockStruct[T []any]() T {
	var v T
	return DefaultMocker.mock(v).(T)
}
