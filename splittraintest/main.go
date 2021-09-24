package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffle[T any](a []T) {
	n := len(a)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func splitTrainTest[T any](a []T, f float64) ([]T, []T) {
	aa := make([]T, len(a), len(a))
	copy(aa, a)
	shuffle(aa)
	i := int(float64(len(aa)) * f)
	return aa[:i], aa[i:]
}

func main() {
	rand.Seed(time.Now().UnixNano())

	vi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	train, test := splitTrainTest(vi, 0.75)
	fmt.Println(train, test)
}
