package main

import (
	"fmt"
)

func ternaryOp[T any](s bool, t func() T, f func() T) T {
	if s {
		return t()
	}
	return f()
}

func main() {
	fmt.Println(ternaryOp(4 < 5, func() string {
		fmt.Println("left evaluated")
		return "less"
	}, func() string {
		fmt.Println("right evaluated")
		return "greater"
	}))
}
