package others

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	//"gopkg.in/yaml.v2"
	//"reflect"
)

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
	Postgres struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		DB       string `json:"db"`
	} `json:"database"`
}

func Func01() {
	jsonConfig := []byte(`{
    "server":{
      "host":"localhost",
      "port":"8080"},
    "database":{
      "host":"localhost",
      "user":"db_user",
      "password":"supersecret",
      "db":"my_db"}
		}`)
	fmt.Println(string(jsonConfig))
	var config Config
	err := json.Unmarshal(jsonConfig, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Config: %+v\n", config)
}

func Func02() {
	config := Config{
		Server: struct {
			Host string `json:"host"`
			Port string `json:"port"`
		}{
			Host: "localhost",
			Port: "8080",
		},
		Postgres: struct {
			Host     string `json:"host"`
			User     string `json:"user"`
			Password string `json:"password"`
			DB       string `json:"db"`
		}{
			Host:     "localhost",
			User:     "db_user",
			Password: "supersecret",
			DB:       "my_db",
		},
	}

	fmt.Printf("%v", config)
}

type Sender struct {
	BankCode string
	Name     string
	Contact  struct {
		Name  string
		Phone string
	}
}

func FuncSender() {

	s := &Sender{
		BankCode: "BC",
		Name:     "NAME",
		Contact: struct {
			Name  string
			Phone string
		}{
			Name:  "Name",
			Phone: "PHONE",
		},
	}

	fmt.Printf("%v", s)
	spew.Dump("%v", s)

}
