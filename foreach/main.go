package main

import (
	"fmt"
)

func eachFunc[T any](a []T, f func(T) bool) {
	for _, e := range a {
		if !f(e) {
			break
		}
	}
}

func main() {
	eachFunc([]int{1, 2, 3, 4, 5, 6}, func(v int) bool {
		fmt.Println(v)
		return v < 4
	})
	eachFunc([]string{"foo", "bar", "", "zoo"}, func(v string) bool {
		fmt.Println(v)
		return v != ""
	})
}
