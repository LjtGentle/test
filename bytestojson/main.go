package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("./ResWhiteListCfg.bytes")
	if err != nil {
		fmt.Println(err)
		return
	}
	var write interface{}
	err = json.Unmarshal(content, &write)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("111")
	fmt.Println(write)
}
