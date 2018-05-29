package flowcontrol

// Defer

import "fmt"

func Flowcontrol12() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
