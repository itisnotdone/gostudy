package moretypes

import "fmt"

func Moretypes15() {
	var s []int

	printSlice15(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice15(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice15(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice15(s)

	// We can add more than one element at a time.
	s = append(s, 5, 6, 7)
	printSlice15(s)

	// We can add more than one element at a time.
	s = append(s, 5, 6, 7)
	printSlice15(s)
}

func printSlice15(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
