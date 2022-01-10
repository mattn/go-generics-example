package main

import (
    "fmt"
)

func add[S ~string | ~[]byte](buf *[]byte, s S) {
	*buf = append(*buf, s...)
}

func main() {
	var buf []byte
	add(&buf, "foo")
	add(&buf, []byte("bar"))

    fmt.Println(string(buf))
}
