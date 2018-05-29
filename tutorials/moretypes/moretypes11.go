package moretypes

// Slice length and capacity

import (
	"fmt"
)

func Moretypes11() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	fmt.Println()

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)
	fmt.Println()

	// Extend its length.
	s = s[:4]
	printSlice(s)
	fmt.Println()

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
	fmt.Println()

	s = s[:4]
	printSlice(s)
	fmt.Println()
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
