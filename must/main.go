package main

import (
	"os"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err.Error())
	}
	return v
}

func Must2[T1 any, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err.Error())
	}
	return v1, v2
}

func Must3[T1 any, T2 any, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err.Error())
	}
	return v1, v2, v3
}

func main() {
	f := Must(os.Create("test.txt"))
	defer os.Remove("test.txt")
	defer f.Close()
	f.Write([]byte("hello"))
}
