package main

import (
	"encoding/json"
	"fmt"
)

func test() {
	testMap := make(map[int]string)
	testMap[1] = "one"
	testMap[2] = "two"
	testMap[3] = "three"
	v, ok := testMap[4]
	if ok && v == "four" {
		fmt.Println("get")
		return
	}
	fmt.Println("not")
}

func main() {
	test9()
}

// 测试指针的range问题
func test9() {
	type People struct {
		ID   int
		Name string
		Age  int
	}
	ps := make([]People, 0, 10)
	psMap := make(map[int]*People)
	psMap[1] = &People{
		ID:   1,
		Name: "111",
		Age:  11,
	}
	psMap[2] = &People{
		ID:   2,
		Name: "222",
		Age:  22,
	}
	psMap[3] = &People{
		ID:   3,
		Name: "333",
		Age:  33,
	}
	for _, p := range psMap {
		ps = append(ps, *p)
	}
	marshal, err := json.Marshal(ps)
	if err != nil {
		return
	}
	fmt.Printf("ps=%+v\n", string(marshal))

	psMap2 := make(map[int]*People)
	for _, v := range ps {
		psMap2[v.ID] = &v
	}
	bytes, err := json.Marshal(psMap2)
	if err != nil {
		return
	}
	fmt.Printf("psMap2=%+v\n", string(bytes))
}

func test8() {
	testMap := make(map[string]int)
	testMap["1"] = 1
	testMap["2"] = 2
	test7(testMap)
	fmt.Println("int test8 testMap=", testMap)
}

func test7(testMap map[string]int) {
	testMap["3"] = 3
	fmt.Println("int test7 testMap=", testMap)
}

func test6() {
	oneMap := make(map[string]*int)
	i1 := 1
	i2 := 2
	i3 := 3
	i4 := 4
	oneMap["one"] = &i1
	oneMap["two"] = &i2
	oneMap["three"] = &i3
	oneMap["four"] = &i4
	tMap := make(map[interface{}]interface{})
	for k, v := range oneMap {
		tMap[k] = v
	}
	fmt.Printf("oneMap=%+v\n", oneMap)
	fmt.Printf("tMap=%+v\n", tMap)
}

func test5() {
	type People struct {
		ID   int
		Name string
		Age  int
	}
	ps := make([]People, 0, 3)
	p1 := People{
		ID:   1,
		Name: "one",
		Age:  11,
	}
	p2 := People{
		ID:   2,
		Name: "two",
		Age:  22,
	}
	p3 := People{
		ID:   3,
		Name: "three",
		Age:  33,
	}
	ps = append(ps, p1)
	ps = append(ps, p2)
	ps = append(ps, p3)
	fmt.Printf("ps=%+v\n", ps)
	psM := make(map[int]*People)
	for _, v := range ps {
		//v1 := v
		//psM[v.ID] = &v1
		fmt.Printf("v point =%p\n", &v)
		fmt.Printf("v value = %+v\n", v)
		psM[v.ID] = &v
	}
	//for i:=0 ;i<len(ps);i ++ {
	//	psM[ps[i].ID] = &ps[i]
	//}
	fmt.Printf("psM=%+v\n", psM)
	for k, v := range psM {
		fmt.Printf("k=%d,v=%+v\n", k, v)
	}
}

func test4() {
	testMap := make(map[string][]int)
	testMap["1"] = []int{1, 2, 3, 4, 5, 6, 7}
	testMap["2"] = []int{11, 12, 13, 14, 15, 16, 17}
	fmt.Printf("testMap=%+v\n", testMap)
	fmt.Println("len=", len(testMap["1"]))
}

// map map
func test3() {
	testMM := make(map[string]map[string]int)
	testMM["one"] = map[string]int{
		"1": 1,
	}
	testMM["two"] = map[string]int{
		"2": 2,
	}
	testMM["two"] = map[string]int{
		"3": 3,
	}
	fmt.Printf("testMM=%+v", testMM)
}

func test2() {
	type stage struct {
		id      int
		version string
		title   string
	}
	stageMap := make(map[int]stage)
	stageMap[1] = stage{
		id:      1,
		version: "11",
		title:   "yyy",
	}
	stageMap[2] = stage{
		id:      2,
		version: "22",
		title:   "xxx",
	}
	s1 := stage{
		id:      3,
		version: "33",
		title:   "ggg",
	}
	stageMap[3] = s1
	fmt.Println(stageMap)
	bytes, err := json.Marshal(stageMap)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println(string(bytes))

	var m map[string]interface{}     //声明map
	m = make(map[string]interface{}) //必须初始化map分配内存
	m["username"] = "user1"
	m["age"] = 18
	m["sex"] = "man"
	fmt.Println(m)
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("json.marshal failed,err:", err)
		return
	}
	fmt.Printf("%s\n", string(data))
}
