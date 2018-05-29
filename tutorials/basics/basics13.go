package basics

// Type conversions

import (
	"fmt"
	"math"
)

func Basics13() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}