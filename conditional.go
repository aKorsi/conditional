package conditional

import "reflect"

func Ternary[T any](con bool, val1, val2 T) T {
	switch con {
	case true:
		return val1
	default:
		return val2
	}
}

func NilCoalescing[T any](val1, val2 T) T {
	switch reflect.ValueOf(val1).IsNil() {
	case true:
		return val2
	default:
		return val1
	}
}

func OptionalChaining[T1, T2 any](val *T1, f func() *T2) *T2 {
	switch reflect.ValueOf(val).IsNil() {
	case true:
		return nil
	default:
		return f()
	}
}
