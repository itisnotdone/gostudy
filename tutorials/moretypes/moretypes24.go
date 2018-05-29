package moretypes

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func blah(fn func(float64, float64) float64, a, b float64) float64 {
	return fn(a, b)
}

func Moretypes24() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(blah(hypot, 7, 8))
	fmt.Println(compute(math.Pow))
}
