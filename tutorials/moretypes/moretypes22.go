package moretypes

import "fmt"

func Moretypes22() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	m["blahblah"] = 33
	v, ok = m["blahblah"]
	fmt.Println("The value:", v, "Present?", ok)
}
