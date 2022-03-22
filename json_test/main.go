package main

import (
	"encoding/json"
	"fmt"
)

func test_map() {
	tmap := make(map[string]interface{})
	str := "{\"id\":\"13\"}"
	err := json.Unmarshal([]byte(str), &tmap)
	if err != nil {
		fmt.Printf("err=%+v\n",err)
		return
	}
	fmt.Printf("tmap=%+v\n",tmap)
	v,ok:= tmap["id"]
	if !ok {
		fmt.Printf("11111\n")
		return
	}
	fmt.Printf("v=%+v\n",v)
}

func test_map2() {
	classMap := make(map[int][]int64)
	classMap[1] = []int64{1,2,3}
	classMap[2] = []int64{1,9,3}
	classMap[3] = []int64{2,9,3}
	b, err := json.Marshal(&classMap)
	if err != nil {
		return
	}
	fmt.Printf("b=%s\n",string(b))
}

func add(classMap map[int]string, classify int,subType string) {
	v, ok := classMap[classify]
	if !ok {
		classMap[classify] = subType
	}else {
		classMap[classify] = v + "," +subType
	}
}

func testSlice() {
	ids := [5]int{1,2,3,4,5}
	is := ids[:]
	b, err := json.Marshal(is)
	if err != nil {

		return
	}
	fmt.Printf("b=%s\n",string(b))
}

func main() {
	//test_map2()
	testSlice()
	//classMap := make(map[int]string)
	//add(classMap,1,"7")
	//add(classMap,1,"8")
	//add(classMap,2,"4")
	//fmt.Printf("classMap=%+v\n",classMap)

}