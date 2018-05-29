package moretypes

import "fmt"

func Moretypes10() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4] // 3 ~ 7
	fmt.Println(s)

	s = []int{2, 3, 5, 7, 11, 13}
	s = s[:2]
	fmt.Println(s)

	s = []int{2, 3, 5, 7, 11, 13}
	s = s[1:]
	fmt.Println(s)
}
