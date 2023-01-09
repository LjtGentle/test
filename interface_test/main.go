package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	test3()
}

// test3 为了测试 str json.Unmarshal ->interface{}，
// interface{}是slice，给slice的对象多加一个字段
func test3() {
	str := "[{\"name\":\"gentle\",\"age\":10}]"
	var i any
	err := json.Unmarshal([]byte(str), &i)
	if err != nil {
		fmt.Println("json unmarshal err=", err)
		return
	}
	fmt.Printf("i=%+v\n", i)
	if reflect.Slice == reflect.TypeOf(i).Kind() {
		fmt.Println("is slice")
	}
}

func test2() {
	is := make([]interface{}, 0, 10)
	i := 10
	is = append(is, i)
	i = 11
	is = append(is, i)
	i = 12
	is = append(is, i)
	fmt.Printf("is type=%T, is=%+v\n", is, is)
	var iss interface{}
	iss = is
	nums, ok := iss.([]int)
	if !ok {
		fmt.Printf("assert failed iss type %T, isss=%+v\n", iss, iss)
		data, err := json.Marshal(iss)
		if err != nil {
			fmt.Println("json marshal err=", err)
			return
		}
		var ns []int
		err = json.Unmarshal(data, &ns)
		if err != nil {
			fmt.Println("json unmarshal err=", err)
			return
		}
		fmt.Printf("ns type %T, ns=%+v\n", ns, ns)

	} else {
		fmt.Println("nums=", nums)
	}

}

func test() {
	// 字符是数字，能否反射成数字类型
	numStr := "18"
	var i interface{}
	i = numStr
	findType(i)
	str, ok := i.(string)
	if !ok {
		fmt.Println("assert failed")
		return
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return
	}
	fmt.Printf("num=%f, T=%T\n", float64(num), float64(num))
	if float64(num) == 18 {
		fmt.Println("ok")
	}

}

func findType(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case uint:
		fmt.Println(x, "is uint")
	case int64:
		fmt.Println(x, "is int64")
	case uint64:
		fmt.Println(x, "is uint64")
	case uint32:
		fmt.Println(x, "is uint32")
	case int32:
		fmt.Println(x, "is int32")
	case int16:
		fmt.Println(x, "is int16")
	case uint16:
		fmt.Println(x, "is uint16")
	case int8:
		fmt.Println(x, "is int8")
	case uint8:
		fmt.Println(x, "is uint8")
	case float64:
		fmt.Println(x, "is float64")
	case float32:
		fmt.Println(x, "is float32")

	default:
		fmt.Println("非数字类型")

	}
}
