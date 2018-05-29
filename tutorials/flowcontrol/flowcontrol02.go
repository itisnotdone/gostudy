package flowcontrol

// For continued, similar to while

import "fmt"

func Flowcontrol02() {
	sum := 1
	for sum < 1000 {
		fmt.Println(sum)
		sum += sum
	}
	fmt.Println(sum)
}
