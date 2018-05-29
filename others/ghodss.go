//http://ghodss.com/2014/the-right-way-to-handle-yaml-in-golang/
package others

import (
	//	"encoding/json"
	"fmt"
	//	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
	//	"reflect"
)

type Person struct {
	Name string `json:"name" yaml:"name"`
	Age  int    `json:"age" yaml:"age"`
	Job  []struct {
		Name   string
		Salary int
	}
}

func DoMarshalYAML() {
	p := Person{Name: "John",
		Age: 30,
		Job: []struct {
			Name   string
			Salary int
		}{
			{
				Name:   "blah",
				Salary: 100,
			},
			{
				Name:   "aaaa",
				Salary: 200,
			},
		},
	}
	y, err := yaml.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(y))

	//p.Job + { Name: "bbbb", Salary: 300, }
	//fmt.Printf("%v", p.Job)
	//y, err = yaml.Marshal(p)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(y))
}

//func DoMarshalJSON() {
//	p := Person{"John", 30}
//	j, err := json.Marshal(p)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(string(j))
//	fmt.Println(reflect.TypeOf(j))
//}
//
//func DoUnmarshalJSON() {
//	j := []byte(`{"name": "John", "age": 30}`)
//	var p Person
//	err := json.Unmarshal(j, &p)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	spew.Dump(p)
//}
