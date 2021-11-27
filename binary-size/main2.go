package main

//go:noinline
func addInt(a, b int) int {
	return a + b
}

//go:noinline
func addString(a, b string) string {
	return a + b
}

func main() {
	println(addInt(1, 2))
	println(addString("foo", "bar"))
}
