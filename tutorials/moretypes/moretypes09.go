package moretypes

import (
	"fmt"
	"reflect"
)

func Moretypes09() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	fmt.Println(reflect.TypeOf(q))

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	fmt.Println(reflect.TypeOf(r))

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}
