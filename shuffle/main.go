package main

import (
    "fmt"
    "math/rand"
)

func shuffle[T any](a []T) {
    n := len(a)
    for i := n - 1; i >= 0; i-- {
        j := rand.Intn(i + 1)
        a[i], a[j] = a[j], a[i]
    }
}

func main() {
    vi := []int{1,2,3,4,5,6,7,8,9,10}
    shuffle(vi)
    fmt.Println(vi)
}

