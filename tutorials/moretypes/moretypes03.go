package moretypes

import "fmt"

type Vertex3 struct {
	X int
	Y int
}

func Moretypes03() {
	v := Vertex3{1, 2}
	fmt.Println(v.X)
	fmt.Println(v.Y)

	v.X = 4
	v.Y = 5
	fmt.Println(v.X)
	fmt.Println(v.Y)
}
