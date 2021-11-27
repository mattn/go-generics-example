package main

import (
	"fmt"
	"reflect"
)

type Addable interface {
	int | int64
}

func AddOne[T Addable](t T) T {
	return t + 1
}

func main() {
	// reflect know only instantiated-types.
	vf := reflect.ValueOf(AddOne[int64])
	fmt.Println(vf.Type())       // func(int64) int64
	fmt.Println(vf.Type().In(0)) // int64
}
