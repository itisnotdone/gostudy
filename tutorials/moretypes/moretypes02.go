package moretypes

import (
	"fmt"
	"reflect"
)

type Vertex2 struct {
	X int
	Y int
}

func Moretypes02() {
	fmt.Println(Vertex2{1, 2})

	vt2 := Vertex2{1, 2}
	fmt.Println(vt2)
	fmt.Println(reflect.TypeOf(vt2))
}
