package main

import (
    "fmt"
)

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

func main() {
    vi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    result := reduceFunc(vi, func(lhs, rhs int) int {
        return lhs + rhs
    }, 1)
    fmt.Println(result)

    vs := []string{"x", "y", "z"}
    s := reduceFunc(vs, func(lhs, rhs string) string {
        return lhs + rhs
    }, "a")
    fmt.Println(s)
}

