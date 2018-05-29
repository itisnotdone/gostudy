package moretypes

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		fmt.Println("sum is", sum)
		return sum
	}
}

func Moretypes25() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println("i is", i)
		fmt.Println(pos(i), neg(-2*i))
		fmt.Println()
	}
}
