package fundamentals

import (
	"fmt"
	"strings"
)

func Func111() {
	module := "function basics"
	author := "don"

	fmt.Println(converter(module, author))
}

func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)

	return module, author
}
