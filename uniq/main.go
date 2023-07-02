package main

import (
	"fmt"
	"math/rand"
	"time"
)

func uniq[T comparable](a []T) []T {
	if len(a) < 2 {
	    return a
	}
	
	u := make([]T, 0, len(a))
	m := make(map[T]bool)
	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = true
			u = append(u, v)
		}
	}
	return u
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(uniq([]int{4, 3, 2, 3, 1, 5, 1}))
}
