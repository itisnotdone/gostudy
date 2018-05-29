package basics

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func Basics6() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
