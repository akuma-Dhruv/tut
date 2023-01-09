package main

import "fmt"

func swap(x, y string) (string, string, int) {
	var a = 10
	return y, x, a
}

func main() {
	a, b, c := swap("hello", "world")
	fmt.Println(a, b, c)
}
