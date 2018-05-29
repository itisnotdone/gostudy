//
//
//
package fundamentals

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
)

var (
// Package level variables are global
// // // name   string
// // // course string
// // name, course string
// // module float64
// name, course, module = "don", "golang fundamentals", 3.2
)

func Hello() {
	name := "don"
	course := "golang fundamentals"
	module := 3.2
	ptr := &module

	fmt.Println("blah", runtime.GOOS)
	fmt.Println("Name is set to", name)
	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is set to", module)
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))

	a := 10.000000
	b := 3

	fmt.Println("\nA is type", reflect.TypeOf(a), "and B is of type", reflect.TypeOf(b))
	c := int(a) + b
	fmt.Println("\nC has value:", c, "and is of type:", reflect.TypeOf(c))

	// points
	fmt.Println("Memory address of *module* variable is", &module)
	// referencing and de-referencing points
	fmt.Println(ptr, reflect.TypeOf(ptr), *ptr, reflect.TypeOf(*ptr))

	fmt.Println("\nHi", name, "you are currently watching", course)
	changeCourse(&course)
	fmt.Println(course)

	//	fmt.Println(os.Environ())
	//	for _, env := range os.Environ() {
	//		fmt.Println(env)
	//	}

	fmt.Println("\n", os.Getenv("LOGNAME"))
}

func changeCourse(course *string) {
	*course = "first look: native docker clustering"
	fmt.Println("trying to change your course to", *course)
	fmt.Println(reflect.TypeOf(*course))
}

//
//
//
