package main

import (
    "fmt"
)

type Addable interface {
	type int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64, complex64, complex128
}

func generator[T Addable](a T) func() T {
    return func() T {
        a = a + 1
        return a
    }
}

func main() {
    f := generator(0)
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())
}

