// NOTE This won't work yet
// https://github.com/golang/go/issues/48609

package main

import (
    "constraints"
    "context"
    "fmt"
)

func makeChan[T constraints.Chan[E], E any](ctx context.Context, arr []E) T {
    ch := make(T)
    go func() {
        defer close(ch)
        for _, v := range arr {
            select {
            case <-ctx.Done():
                return
            default:
            }
            ch <- v
        }
    }()
    return ch
}

func main() {
    for v := range makeChan(context.Background(), []int{1, 2, 3}) {
        fmt.Println(v)
    }
}
