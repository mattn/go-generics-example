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
		fmt.Println("left excuted")
		return "less"
	}, func() string {
		fmt.Println("right excuted")
		return "greater"
	}))
}
