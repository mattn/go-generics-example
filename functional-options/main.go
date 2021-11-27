package main

import "fmt"

type Addable interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

type Foo[T Addable] struct{ v T }

func (foo *Foo[T]) Add(v T) {
	foo.v += v
}

type Option[T Addable] func(*Foo[T])

func WithValue[T Addable](v T) Option[T] {
	return func(f *Foo[T]) { f.v = v }
}

func NewFoo[T Addable](options ...Option[T]) *Foo[T] {
	foo := new(Foo[T])
	for _, opt := range options {
		opt(foo)
	}
	return foo
}

func main() {
	foo := NewFoo(WithValue(3.14))
	foo.Add(4)
	// 7.140000000000001 (float64)
	fmt.Printf("%v (%T)\n", foo.v, foo.v)
}
