package main

import (
	"fmt"
)

type comparable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

func max[T comparable](a []T) T {
	m := a[0]
	for _, v := range a {
		if m < v {
			m = v
		}
	}
	return m
}

func min[T comparable](a []T) T {
	m := a[0]
	for _, v := range a {
		if m > v {
			m = v
		}
	}
	return m
}

func main() {
	vi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := max(vi)
	fmt.Println(result)

	vi = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result = min(vi)
	fmt.Println(result)
}
