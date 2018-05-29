package moretypes

import "fmt"

type Vertex5 struct {
	X, Y int
}

var (
	v1 = Vertex5{1, 2}         // has type Vertex
	v2 = Vertex5{X: 1}         // Y:0 is implicit
	v3 = Vertex5{}             // X:0 and Y:0
	v4 = Vertex5{X: 1, Y: 100} // Y:0 is implicit

	p = &Vertex5{1, 2} // has type *Vertex
)

func Moretypes05() {
	fmt.Println(v1, v2, v3, v4)
	fmt.Println(p, *p, p.X, (*p).X)
}
