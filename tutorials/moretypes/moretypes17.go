package moretypes

/*
Range continued
You can skip the index or value by assigning to _.

If you only want the index, drop the ", value" entirely.
*/

import "fmt"

func Moretypes17() {

	pow := make([]int, 10)
	fmt.Println(pow, len(pow), cap(pow))

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}

}
