package main

import "fmt"

func makeComparableFunc[T comparable]() func(lhs, rhs T) bool {
	return func(lhs, rhs T) bool {
		return lhs == rhs
	}
}

func main() {
	// https://github.com/golang/go/issues/39632
	//equal := func[T comparable](lhs, rhs T) bool {
	//    return lhs == rhs
	//}

	equal := makeComparableFunc[int]()
	fmt.Println(equal(1, 2))
}
