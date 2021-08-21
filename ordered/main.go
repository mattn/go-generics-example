// NOTE This won't work yet
// https://github.com/golang/go/issues/45458

package main

import (
    "fmt"
    "constraints"
)

func sortFunc[T constraint.Ordered](a []T) {
	for i := 0; i < len(a) - 1; i++ {
		for j := 0; j < len(a) - i - 1; j++ {
			if a[j] > a[j + 1] {
				a[j], a[j + 1] = a[j + 1], a[j]
			}
		}
	}
}

func main() {
    vi := []int{1,2,3,4,5,6}
    sortFunc(vi)
    fmt.Println(vi)
}
