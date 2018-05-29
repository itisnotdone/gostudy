package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func Basics1() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("\nMy favorite number is", rand.Intn(100))
}
