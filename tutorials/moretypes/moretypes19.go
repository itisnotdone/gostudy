package moretypes

import "fmt"

type Vertex19 struct {
	Lat, Long float64
}

var m map[string]Vertex19

func Moretypes19() {

	m = make(map[string]Vertex19)

	m["Bell Labs"] = Vertex19{
		40.68433, -74.39967,
	}

	m["Don"] = Vertex19{
		40, -74,
	}

	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Don"])

	nm := make(map[interface{}]interface{})
	nm["Don"] = Vertex19{
		40, -74,
	}
	fmt.Println(m["Don"])

}
