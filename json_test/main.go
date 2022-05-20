package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func test_map() {
	tmap := make(map[string]interface{})
	str := "{\"id\":\"13\"}"
	err := json.Unmarshal([]byte(str), &tmap)
	if err != nil {
		fmt.Printf("err=%+v\n", err)
		return
	}
	fmt.Printf("tmap=%+v\n", tmap)
	v, ok := tmap["id"]
	if !ok {
		fmt.Printf("11111\n")
		return
	}
	fmt.Printf("v=%+v\n", v)
}

func test_map2() {
	classMap := make(map[int][]int64)
	classMap[1] = []int64{1, 2, 3}
	classMap[2] = []int64{1, 9, 3}
	classMap[3] = []int64{2, 9, 3}
	b, err := json.Marshal(&classMap)
	if err != nil {
		return
	}
	fmt.Printf("b=%s\n", string(b))
}

func add(classMap map[int]string, classify int, subType string) {
	v, ok := classMap[classify]
	if !ok {
		classMap[classify] = subType
	} else {
		classMap[classify] = v + "," + subType
	}
}

func testSlice() {
	ids := [5]int{1, 2, 3, 4, 5}
	is := ids[:]
	b, err := json.Marshal(is)
	if err != nil {

		return
	}
	fmt.Printf("b=%s\n", string(b))
}

func main() {
	//test_map2()
	//testSlice()
	test4()
	//classMap := make(map[int]string)
	//add(classMap,1,"7")
	//add(classMap,1,"8")
	//add(classMap,2,"4")
	//fmt_test.Printf("classMap=%+v\n",classMap)

}
type PageInfo struct {
	Filter   uint32 `json:"filter"`
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"page_size"`
	Offset   uint32 `json:"-"`
}

func test4() {
	var p PageInfo
	str := "{\"filter\":1,\"page\":2,\"page_size\":20}"
	if err := json.Unmarshal([]byte(str), &p); err != nil {
		 fmt.Printf("错误的分页数据:%v\n", str)
		return
	}
	fmt.Printf("p=%+v",p)
}


// 删除某个下标
func test3() {
	is := []int{1,2,3,4,5,6,7}
	intSlice := is[:]
	fmt.Println("intSlice=",intSlice)
	for k,v := range intSlice {
		fmt.Printf("k=%d,v=%d\n",k,v)
	}
	// 删除下标为3的
	index := 3
	pre := intSlice[:index]
	tail := intSlice[index+1:]
	res := make([]int,0,10)
	res = append(pre,tail...)
	fmt.Println("res=",res)


}


// json 转对象

func test2() {
	//str := "{\"primary\":{\"name\":\"重拳\",\"price\":452,\"icon\":\"https://wuji-1254960240.file.myqcloud.com/xy/ingame_sy/web4ea55d73-3374-4a63-bcc4-0aae6c0bf31c.png\",\"desc\":\"\",\"attributes\":[{\"attribute\":\"被动-专精（张飞）效果调整为：\",\"desc\":\"【守护机关】可以对目标区域敌人造成伤害，并附带50%减速持续1s，但无法再为队友提供护盾，【崩裂践踏】可额外对目标造成50%减速，持续1s\"}]}}"
	str := "[{\\\"name\\\":\\\"被动-专精（张飞）效果调整为：\\\",\\\"desc\\\":\\\"【守护机关】可以对目标区域敌人造成伤害，并附带50%减速持续1s，但无法再为队友提供护盾，【崩裂践踏】可额外对目标造成50%减速，持续1s\\\"},{\\\"name\\\":\\\"被动-力量进阶效果调整为：\\\",\\\"desc\\\":\\\"额外物理攻击达到100时，释放【守护机关】和【崩裂践踏】后，会持续对周围造成法术灼烧伤害，并降低受到伤害的目标30%的输出\\\"}]"
	str = strings.Replace(str,"\\","",-1)
	fmt.Println("str=",str)
	var obj interface{}
	err := json.Unmarshal([]byte(str), &obj)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("obj=", obj)
	b, err := json.Marshal(obj)
	if err != nil {
		return
	}
	fmt.Println("b=",string(b))

}
