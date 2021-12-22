package main

type Int64 interface {
	~int64 | ~uint64
}

type NumberPlayer[T Int64] interface {
	AddCounter() T
	MinusCounter() T
}

type Counter[T Int64] struct {
	t T
	counter T
}

func (c *Counter[T]) AddCounter() T {
	c.t += c.counter
	return c.t
}

func (c *Counter[T]) MinusCounter() T {
	c.t -= c.counter
	return c.t
}

func NewNumberPlayer[T Int64](t T, counter T) NumberPlayer[T] {
	return &Counter[T]{
		t: t,
		counter: counter,
	}
}

func main() {
	numbers := NewNumberPlayer(int64(32), int64(2))
	println(numbers.AddCounter())
	println(numbers.MinusCounter())

	numbersAgain := NewNumberPlayer(uint64(32), uint64(9))
	println(numbersAgain.AddCounter())
	println(numbersAgain.MinusCounter())
}
