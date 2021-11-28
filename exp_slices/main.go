package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	a := []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
		"END",
	}
	a = a[:12]
	fmt.Println(cap(a)) // 13
	a = slices.Clip(a)
	fmt.Println(cap(a)) // 12

	a = []string{
		"foo",
		"foo",
		"bar",
		"baz",
		"foo",
		"baz",
	}
	a = slices.Compact(a)
	fmt.Println(a) // [foo bar baz foo baz]

	fmt.Println(slices.Contains(a, "zoo")) // false

	a = slices.Delete(a, 1, 3)
	fmt.Println(a) // [foo foo baz]

	fmt.Println(slices.Index(a, "baz")) // 2

	fmt.Println(cap(a)) // 6
	a = slices.Grow(a, 5)
	fmt.Println(cap(a)) // 12 WTF?

	a = slices.Insert(a, 0, "noo")
	fmt.Println(a) // [noo foo foo baz]
}
