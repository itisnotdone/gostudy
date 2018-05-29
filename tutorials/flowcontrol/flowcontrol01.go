package flowcontrol

import "fmt"

func Flowcontrol01() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(i, sum)
	}
	fmt.Println(sum)
}
