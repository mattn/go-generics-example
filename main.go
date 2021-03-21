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

func reduceFunc[T any](a []T, f func(T, T) T, initial interface{}) T {
    if len(a) == 0 || f == nil {
        var vv T
        return vv
    }

    l := len(a) - 1

    reduce := func(a []T, ff func(T, T) T, memo T, startPoint, direction, length int) T {
        result := memo
        index := startPoint

        for i := 0; i <= length; i++ {
            result = ff(result, a[index])
            index += direction
        }

        return result
    }

    if initial == nil {
        return reduce(a, f, a[0], 1, 1, l-1)
    }

    return reduce(a, f, initial.(T), 0, 1, l)
}

type comparable interface {
	type int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64
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

    // using reduceFunc
    vi = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    result := reduceFunc(vi, func(lhs, rhs int) int {
        return lhs + rhs
    }, 1)
    fmt.Println(result)

    // using max
    vi = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    result = max(vi)
    fmt.Println(result)

    // using min
    vi = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    result = min(vi)
    fmt.Println(result)
}
