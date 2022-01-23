package main

import (
	"testing"
)

func addInt(a, b int) int {
	return a + b
}

func addString(a, b string) string {
	return a + b
}

func BenchmarkWithoutGenerics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = addInt(1, 2)
		_ = addString("foo", "bar")
	}
}

type Addable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | complex64 | complex128 | string
}

func add[T Addable](a, b T) T {
	return a + b
}

func BenchmarkWithGenerics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = add(1, 2)
		_ = add("foo", "bar")
	}
}

func addInterface(a, b interface{}) interface{} {
	switch v := a.(type) {
	case int:
		return v + b.(int)
	case string:
		return v + b.(string)
	default:
		panic("unrecognized type")
	}
}

func BenchmarkWithInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = addInterface(1, 2)
		_ = addInterface("foo", "bar")
	}
}
