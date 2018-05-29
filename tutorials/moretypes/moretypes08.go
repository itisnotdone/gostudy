package moretypes

/* A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes. */

import (
	"fmt"
	"reflect"
)

func Moretypes08() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	fmt.Println(reflect.TypeOf(names))
	fmt.Println()

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
