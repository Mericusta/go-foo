package mockfoo

import (
	"fmt"
	"math"
	"reflect"
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

func toInterface(v interface{}) interface{} {
	return v
}

func mockSingleType[T any](iv interface{}) (T, bool) {
	switch iv.(type) {
	case int:
		iv = int(math.MaxInt)
	case string:
		iv = string(mockComplexType[[]rune]([]rune{}))
	default:
		return iv.(T), false
	}
	return iv.(T), true
}

func mockComplexType[T any](iv interface{}) T {
	typeI := reflect.TypeOf(iv)
	fmt.Printf("typeI = %v\n", typeI)
	switch typeI.Kind() {
	case reflect.Array:
		iv = reflect.MakeSlice(typeI, 8, 8).Interface()
	case reflect.Slice:
		iv = reflect.MakeSlice(typeI, 0, 8).Interface()
	case reflect.Map:
		iv = reflect.MakeMap(typeI).Interface()
	case reflect.Struct:
		numField := typeI.NumField()
		for i := 0; i < numField; i++ {

		}
		// structValue := reflect.New(typeI)
		// structValue.
	}
	return iv.(T)
}

func Mock[T any](v T) T {
	mockValue, isSingle := mockSingleType[T](v)
	if isSingle {
		return mockValue
	}
	return mockComplexType[T](v)
}

func TMock[T any]() T {
	var v T
	return Mock(v)
}

func MockFoo() {
	// mock by value
	// var i1 int
	var i1 []int
	i1 = Mock(i1)
	fmt.Printf("i1 = %v\n", i1)

	var s1 string
	s1 = Mock(s1)
	fmt.Printf("s1 = %v\n", s1)

	// // mock by type
	// i2 := TMock[int]()
	// fmt.Printf("i2 = %v\n", i2)

	var (
	// b     bool
	// i     int
	// i8    int8
	// i16   int16
	// i32   int32
	// i64   int64
	// ui    uint
	// ui8   uint8
	// ui16  uint16
	// ui32  uint32
	// ui64  uint64
	// f32   float32
	// f64   float64
	// str   string
	// s  []int
	// ss [][]int
	// sm []map[int]string
	// m  map[int]int
	// mm map[int]map[int]string
	// ms map[int][]string
	)

	// b = Mock[bool]()
	// fmt.Printf("b = %v\n", b)
	// i = Mock[int]()
	// fmt.Printf("i = %v\n", i)
	// i8 = Mock[int8]()
	// fmt.Printf("i8 = %v\n", i8)
	// i16 = Mock[int16]()
	// fmt.Printf("i16 = %v\n", i16)
	// i32 = Mock[int32]()
	// fmt.Printf("i32 = %v\n", i32)
	// i64 = Mock[int64]()
	// fmt.Printf("i64 = %v\n", i64)
	// ui = Mock[uint]()
	// fmt.Printf("ui = %v\n", ui)
	// ui8 = Mock[uint8]()
	// fmt.Printf("ui8 = %v\n", ui8)
	// ui16 = Mock[uint16]()
	// fmt.Printf("ui16 = %v\n", ui16)
	// ui32 = Mock[uint32]()
	// fmt.Printf("ui32 = %v\n", ui32)
	// ui64 = Mock[uint64]()
	// fmt.Printf("ui64 = %v\n", ui64)
	// f32 = Mock[float32]()
	// fmt.Printf("f32 = %v\n", f32)
	// f64 = Mock[float64]()
	// fmt.Printf("f64 = %v\n", f64)
	// str = Mock[string]()
	// fmt.Printf("s = %v\n", str)

	// s = MockSlice(s)
	// fmt.Printf("slice = %v\n", s)
	// ss = MockSlice(ss)
	// fmt.Printf("ss = %v\n", ss)
	// sm = MockSlice(sm)
	// fmt.Printf("sm = %v\n", sm)
	// m = Mock[map[int]int]()
	// fmt.Printf("m = %v\n", m)
	// ms = Mock[map[int][]string]()
	// fmt.Printf("ms = %v\n", ms)
}

// type Mocker struct {
// 	// IntMocker
// }

// var DefaultMocker = &Mocker{}

// func mockSingle(v interface{}) (interface{}, bool) {
// 	switch v.(type) {
// 	case bool:
// 		return rand.Intn(2) == 0, true
// 	case int:
// 		return int(math.MaxInt), true
// 	case int8:
// 		return int8(math.MaxInt8), true
// 	case int16:
// 		return int16(math.MaxInt16), true
// 	case int32:
// 		return int32(math.MaxInt32), true
// 	case int64:
// 		return int64(math.MaxInt64), true
// 	case uint:
// 		return uint(math.MaxUint), true
// 	case uint8:
// 		return uint8(math.MaxUint8), true
// 	case uint16:
// 		return uint16(math.MaxUint16), true
// 	case uint32:
// 		return uint32(math.MaxUint32), true
// 	case uint64:
// 		return uint64(math.MaxUint64), true
// 	case float32:
// 		return float32(math.MaxFloat32), true
// 	case float64:
// 		return float64(math.MaxFloat64), true
// 	case string:
// 		return string("string"), true
// 	default:
// 		return v, false
// 	}
// }

// func mockSingleT[T any]() (T, bool) {
// 	var v T
// 	iv, ok := mockSingle(v)
// 	if !ok {
// 		return v, false
// 	}
// 	return iv.(T), true
// }

// func mockSlice(v interface{}) (interface{}, bool) {
// 	bs, ok := v.([]bool)
// 	if ok {
// 		return bs, true
// 	}
// 	is, ok := v.([]int)
// 	if ok {
// 		return is, true
// 	}
// 	i8s, ok := v.([]int8)
// 	if ok {
// 		return i8s, true
// 	}
// 	i16s, ok := v.([]int16)
// 	if ok {
// 		return i16s, true
// 	}
// 	i32s, ok := v.([]int32)
// 	if ok {
// 		return i32s, true
// 	}
// 	i64s, ok := v.([]int64)
// 	if ok {
// 		return i64s, true
// 	}
// 	uis, ok := v.([]uint)
// 	if ok {
// 		return uis, true
// 	}
// 	ui8s, ok := v.([]uint8)
// 	if ok {
// 		return ui8s, true
// 	}
// 	ui16s, ok := v.([]uint16)
// 	if ok {
// 		return ui16s, true
// 	}
// 	ui32s, ok := v.([]uint32)
// 	if ok {
// 		return ui32s, true
// 	}
// 	ui64s, ok := v.([]uint64)
// 	if ok {
// 		return ui64s, true
// 	}
// 	f32s, ok := v.([]float32)
// 	if ok {
// 		return f32s, true
// 	}
// 	f64s, ok := v.([]float64)
// 	if ok {
// 		return f64s, true
// 	}
// 	strs, ok := v.([]string)
// 	if ok {
// 		return strs, true
// 	}
// 	return nil, false
// }

// func mockSliceT[T any](ts []T) ([]T, bool) {

// 	// mockSlice([])

// 	return nil, false
// }

// func traitSlice[T any, ST []T](s ST) (T, ST) {
// 	var e T
// 	return e, s
// }

// func traitMap[K comparable, V any, MT map[K]V](m MT) (K, V, MT) {
// 	var (
// 		k K
// 		v V
// 	)
// 	return k, v, m
// }

// func trait(v interface{}) (interface{}, bool) {
// 	return nil, false
// }

// func Mock[T any]() T {
// 	var v T
// 	// return mock[T](v)
// 	trait(v)
// 	// v = iv.(T)
// 	return v
// }

// func MockSlice[T any, ST []T](vs ST) ST {
// 	e, v := traitSlice(vs)
// 	fmt.Printf("e = %v\n", e)
// 	fmt.Printf("v = %v\n", v)

// 	ie, ok := mockSingle(e)
// 	if ok {
// 		vs = make(ST, 0, 8)
// 		for i := 0; i < 8; i++ {
// 			ie, ok := mockSingle(e)
// 			if ok {
// 				vs = append(vs, ie.(T))
// 			}
// 		}
// 	} else {
// 		mockSlice(ie)
// 	}
// 	return vs
// }
