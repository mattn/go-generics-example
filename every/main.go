package main

import (
	"fmt"
)

func every[T any](a []T, f func(T) bool) bool {
	for _, e := range a {
		if !f(e) {
			return false
		}
	}

	return true
}

func main() {
	allEven := every([]int{2, 4, 6, 8, 10}, func(v int) bool {
		return v%2 == 0
	})
	if allEven {
		fmt.Println("All even!")
	}

	allEven = every([]int{2, 3, 4, 5, 6}, func(v int) bool {
		return v%2 == 0
	})
	if !allEven {
		fmt.Println("Some odd")
	}
}
