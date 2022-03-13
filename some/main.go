package main

import (
	"fmt"
)

func some[T any](a []T, f func(T) bool) bool {
	for _, e := range a {
		if f(e) {
			return true
		}
	}

	return false
}

func main() {
	someEven := some([]int{2, 3, 4, 5, 6}, func(v int) bool {
		return v%2 == 0
	})
	if someEven {
		fmt.Println("Some odd")
	}

	someEven = some([]int{1, 3, 5, 7}, func(v int) bool {
		return v%2 == 0
	})
	if !someEven {
		fmt.Println("Only odd")
	}
}
