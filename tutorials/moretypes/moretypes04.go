package moretypes

// Pointers to structs

/*
To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
*/

import (
	"fmt"
	"reflect"
)

type Vertex4 struct {
	X int
	Y int
}

func Moretypes04() {
	v := Vertex4{1, 2}
	p := &v
	fmt.Println((*p).X)
	p.X = 1e9
	fmt.Println(v.X)
	fmt.Println(reflect.TypeOf(v.X))
}
