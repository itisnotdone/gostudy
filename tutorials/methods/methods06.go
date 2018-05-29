package methods

import "fmt"

type Vertex6 struct {
	X, Y float64
}

func (v *Vertex6) Scale6(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex6, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Methods06() {
	v := Vertex6{3, 4}
	v.Scale6(2)
	ScaleFunc(&v, 10)

	p := &Vertex6{4, 3}
	p.Scale6(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
