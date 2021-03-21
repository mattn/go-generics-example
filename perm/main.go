package main

import (
    "fmt"
)

func permInternal[T any](a []T, f func([]T), i int) {
    if i > len(a) {
        f(a)
        return
    }
    permInternal(a, f, i+1)
    for j := i + 1; j < len(a); j++ {
        a[i], a[j] = a[j], a[i]
        permInternal(a, f, i+1)
        a[i], a[j] = a[j], a[i]
    }
}

func perm[T any](a []T, f func([]T)) {
	permInternal(a, f, 0)

}

func main() {
    vi := []int{1, 2, 3}
    perm(vi, func(a []int) {
        fmt.Println(a)
    })
}

