package methods

//
//import (
//	"fmt"
//	"math"
//)
//
//type Abser interface {
//	Abs() float64
//}
//
//func Methods09() {
//	var a Abser
//	f := MyFloat9(-math.Sqrt2)
//	v := Vertex9{3, 4}
//
//	a = f  // a MyFloat9 implements Abser
//	a = &v // a *Vertex9 implements Abser
//
//	// In the following line, v is a Vertex9 (not *Vertex9)
//	// and does NOT implement Abser.
//	a = v
//
//	fmt.Println(a.Abs())
//}
//
//type MyFloat9 float64
//
//func (f MyFloat9) Abs() float64 {
//	if f < 0 {
//		return float64(-f)
//	}
//	return float64(f)
//}
//
//type Vertex9 struct {
//	X, Y float64
//}
//
//func (v *Vertex9) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
