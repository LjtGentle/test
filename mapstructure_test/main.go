package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type People struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {
	input := map[string]interface{}{
		"name": "gentle",
		"age":  "24",
	}
	p := People{}
	err := mapstructure.Decode(input, &p)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("p=", p)

	testMap := map[string]interface{}{
		"name": "",
		"age":  "",
		"addr": "dsfsdf",
	}
	err = mapstructure.Decode(p, &testMap)
	if err != nil {
		fmt.Println("err222=", err)
		return
	}
	fmt.Println("testMap=", testMap)
}
