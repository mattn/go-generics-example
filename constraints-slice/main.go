package main

import (
	"constraints"
	"fmt"
)

func sort[T constraints.Slice[E], E constraints.Ordered](arr T) T {
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	fmt.Println(sort([]int{3,1,2}))
	fmt.Println(sort([]float64{6.2,7.91,2.1}))
}
