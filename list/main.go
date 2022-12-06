package main

import (
	"fmt"
)

type List[T comparable] struct {
	l []T
}

func (l *List[T]) Push(v T) {
	l.l = append(l.l, v)
}

func (l *List[T]) Insert(v T) {
	l.InsertAt(0, v)
	return
}

func (l *List[T]) InsertAt(pos int, v T) {
	l.l = append(l.l[:pos+1], l.l[pos:]...)
	l.l[pos] = v
	return
}

func (l *List[T]) Remove(i int) {
	l.l = l.l[:i+copy(l.l[i:], l.l[i+1:])]
	return
}

func (l *List[T]) Equals(rhs *List[T]) bool {
	if len(l.l) != len(rhs.l) {
		return false
	}
	for i := 0; i < len(l.l); i++ {
		if l.l[i] != rhs.l[i] {
			return false
		}
	}
	return true
}

func (l *List[T]) Clone() *List[T] {
	ll := &List[T]{l: make([]T, len(l.l))}
	copy(ll.l, l.l)
	return ll
}

func (l *List[T]) Slice() []T {
	return l.l
}

func NewList[T comparable](t []T) *List[T] {
	return &List[T]{l: t}
}

func main() {
	l := NewList([]int{1, 2, 3})
	l.Push(4)
	l.Push(5)
	l.Push(6)
	fmt.Println(l.Slice())
	l.Remove(3)
	fmt.Println(l.Slice())
	c := l.Clone()
	fmt.Println(c.Equals(l))
}
