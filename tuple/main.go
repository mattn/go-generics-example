package main

import (
	"fmt"
	"math/rand"
)

type tuple2[T1, T2 any] struct {
	V1 T1
	V2 T2
}

func zip2[T1, T2 any, tuple tuple2[T1, T2]](v1 []T1, v2 []T2) []tuple {
	if len(v1) != len(v2) {
		panic("length of v1 and v2 must be same")
	}
	vv := make([]tuple, len(v1), cap(v1))
	for i := range v1 {
		vv[i].V1 = v1[i]
		vv[i].V2 = v2[i]
	}
	return vv
}

func unzip2[T1, T2 any, tuple tuple2[T1, T2]](vv []tuple) ([]T1, []T2) {
	v1 := make([]T1, len(vv))
	v2 := make([]T2, len(vv))
	for i, v := range vv {
		v1[i] = v.V1
		v2[i] = v.V2
	}
	return v1, v2
}

func shuffle[T any](a []T) []T {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	a, b = unzip2(shuffle(zip2(a, b)))
	fmt.Println(a)
	fmt.Println(b)
}
