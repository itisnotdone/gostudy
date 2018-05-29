package methods

// Methods are functions

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

func Abs(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Methods02() {
	v := Vertex2{3, 4}
	fmt.Println(Abs(v))
}
