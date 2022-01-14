package main

import (
	"constraints"
	"fmt"
)

// type v[T constraints.Ordered] T

type Vector[T constraints.Ordered] struct {
	x, y T
}

func (v *Vector[T]) Add(x, y T) {
	v.x += T(x)
	v.y += T(y)
}

func (v *Vector[T]) String() string {
	return fmt.Sprintf("{x: %v, y: %v}", v.x, v.y)
}

func NewVector[T constraints.Ordered](x, y T) *Vector[T] {
	return &Vector[T]{x: x, y: y}
}

func main() {
	v := NewVector[float64](1, 2)
	v.Add(2, 3)
	fmt.Println(v)
}
