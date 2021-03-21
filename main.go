package main

import (
    "fmt"
)

func filterFunc[T any](a []T, f func(T) bool) []T {
    var n []T
    for _, e := range a {
        if f(e) {
            n = append(n, e)
        }
    }
    return n
}

func mapFunc[T any, M any](a []T, f func(T) M) []M {
    n := make([]M, len(a), cap(a))
    for i, e := range a {
        n[i] = f(e)
    }
    return n
}

func main() {
    vi := []int{1,2,3,4,5,6}

    vi = filterFunc(vi, func(v int) bool {
        return v < 4
    })
    fmt.Println(vi)

    vs := mapFunc(vi, func(v int) string {
        return "<" + fmt.Sprint(v) + ">"
    })
    fmt.Println(vs)
}
