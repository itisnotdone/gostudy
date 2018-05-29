package flowcontrol

import (
	"fmt"
	"time"
)

// Weekday starts from 0, which is Sunday
func Flowcontrol10() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Go work. It's too far away.")
	}
}
