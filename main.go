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

func ternaryOp[T any](s bool, t T, f T) T {
    if s {
        return t
    }
    return f
}

func main() {
    // using filterFunc
    vi := []int{1,2,3,4,5,6}
    vi = filterFunc(vi, func(v int) bool {
        return v < 4
    })
    fmt.Println(vi)

    // using mapFunc
    vs := mapFunc(vi, func(v int) string {
        return "<" + fmt.Sprint(v) + ">"
    })
    fmt.Println(vs)

    // using ternaryOp
    fmt.Println(ternaryOp(len(vs) < 5, "short", "long"))
}
