package main

import (
	"fmt"
)

type Addable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | complex64 | complex128 | string
}

func add[T Addable](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add(1, 2))

	fmt.Println(add("foo", "bar"))
}
