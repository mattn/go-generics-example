package main

import (
    "fmt"
)

func ternaryOp[T any](s bool, t T, f T) T {
    if s {
        return t
    }
    return f
}

func main() {
    fmt.Println(ternaryOp(4 < 5, "less", "greater"))
}

