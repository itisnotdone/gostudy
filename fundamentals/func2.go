package fundamentals

import "fmt"

func Func222() {
	bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 16)
	fmt.Println(bestFinish)
}

func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}
