package moretypes

import "fmt"

func Moretypes07() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[0:6]
	fmt.Println(s)
}
