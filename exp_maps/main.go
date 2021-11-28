package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func main() {
	m := map[int]string{
		123: "foo",
		456: "bar",
	}
	fmt.Println(maps.Keys(m))   // [123 456]
	fmt.Println(maps.Values(m)) // [foo bar]
	m2 := maps.Clone(m)
	fmt.Println(maps.Equal(m2, m)) // true
	m3 := map[int]string{
		789: "baz",
	}
	maps.Copy(m3, m)
	fmt.Println(maps.Equal(m3, m)) // false
}
