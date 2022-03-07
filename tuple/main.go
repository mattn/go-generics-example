package main

import (
	"fmt"
	"math/rand"
)

type tuple2[T1, T2 any] struct {
	V1 T1
	V2 T2
}

func zip2[T1, T2 any](v1 []T1, v2 []T2) []tuple2[T1, T2] {
	vv := make([]tuple2[T1, T2], len(v1))
	for i := range v1 {
		vv[i].V1 = v1[i]
		vv[i].V2 = v2[i]
	}
	return vv
}

func unzip2[T1, T2 any](vv []tuple2[T1, T2]) ([]T1, []T2) {
	v1 := make([]T1, len(vv))
	v2 := make([]T2, len(vv))
	for i, v := range vv {
		v1[i] = v.V1
		v2[i] = v.V2
	}
	return v1, v2
}

func shuffle[T any](a []T) []T {
	n := len(a)
	v := make([]T, len(a))
	copy(v, a)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		v[i], v[j] = v[j], v[i]
	}
	return v
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	a, b = unzip2(shuffle(zip2(a, b)))
	fmt.Println(a)
	fmt.Println(b)
}
