package builtinfoo

import "reflect"

func DeepEqualFoo[T, V any](t T, v V) bool {
	return reflect.DeepEqual(t, v)
}
