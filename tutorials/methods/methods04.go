package methods

// Pointer receivers, to be able to change original value of variables

import (
	"fmt"
	"math"
)

type Vertex4 struct {
	X, Y float64
}

func (v Vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex4) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Methods04() {
	v := Vertex4{3, 4}
	fmt.Println(v.X, v.Y)

	v.Scale(10)
	fmt.Println(v.X, v.Y)

	fmt.Println(v.Abs())
}
