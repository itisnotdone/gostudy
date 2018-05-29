package methods

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Methods01() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	vv := Vertex{10, 11}
	fmt.Println(vv.Abs())
}
