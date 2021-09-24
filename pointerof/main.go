package main

import (
	"fmt"
)

func pointerOf[T any](v T) *T {
	return &v
}

func main() {
	sp := pointerOf("foo")
	fmt.Println(*sp)

	ip := pointerOf(123)
	fmt.Println(*ip)
	*ip = 234
	fmt.Println(*ip)
}
