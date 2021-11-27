package main

type Int64OrFloat64 interface {
	int64 | float64
}

type Int64LikeOrFloat64Like interface {
	~int64 | ~float64
}

func V1[T Int64OrFloat64](v T)         { println(v) }
func V2[T Int64LikeOrFloat64Like](v T) { println(v) }

func main() {
	// V1(1) // NG since 1 is not int64
	V1(int64(1)) // OK since it is int64
	type Int64 int64
	var v Int64 = 1
	// V1(v) // int64 but error since underlying type
	V2(v) // OK since V2 accept underlying type
}
