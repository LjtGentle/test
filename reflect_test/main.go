package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	test03()
}

// test3 为了测试 str json.Unmarshal ->interface{}，
// interface{}是slice，给slice的对象多加一个字段
func test03() {
	str := "[{\"name\":\"gentle\",\"age\":10}]"
	var i interface{}
	err := json.Unmarshal([]byte(str), &i)
	if err != nil {
		fmt.Println("json unmarshal err=", err)
		return
	}
	fmt.Printf("i=%+v\n", i)
	if reflect.Slice == reflect.TypeOf(i).Kind() {
		fmt.Println("is slice")
	}
	v := reflect.ValueOf(i)
	len := v.Len()
	for j := 0; j < len; j++ {
		if v.Index(j).Elem().Kind() == reflect.Map {
			fmt.Println("is map")
		}

		var key = "key"
		var value = "value"
		var name = "name"
		iv := v.Index(j).Elem().MapIndex(reflect.ValueOf(name))
		fmt.Printf("iv.Kind()=%+v\n", iv.Kind())
		IV := iv.Interface()
		str, ok := IV.(string)
		if !ok {
			fmt.Println("断言失败")
			return
		}
		fmt.Println("str=", str)
		fmt.Printf("iv.String()=%+v\n", iv.String())
		fmt.Printf("MapIndex=%+v\n", iv)
		v.Index(j).Elem().SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
		fmt.Printf("v.Index=%+v\n", v.Index(j))
	}
	fmt.Println("i=", i)
}

type Person struct {
	Age  int    `json:"age"`
	Addr string `json:"addr"`
	Name string `json:"name"`
}

func (p Person) Hello(name string) {
	fmt.Printf("%s say: hello %s", name, p.Name)
}

func test02() {
	//var o interface{}
	//p := Person{
	//	Age:  18,
	//	Addr: "gz",
	//	Name: "wawa",
	//}
	//o = &p
	//o = p
	//t := reflect.TypeOf(o)
	//fmt.Println("t=", t)
	//fmt.Println("t.Name()=", t.Name())
	//
	//v := reflect.ValueOf(o)
	//fmt.Println("v=", v)
	////v = v.Elem()
	//for i := 0; i < t.NumField(); i++ {
	//	f := t.Field(i)
	//	fmt.Println(f.Name, "---", f.Type)
	//	val := v.Field(i).Interface()
	//	fmt.Println("val=", val)
	//}
	//o = &p
	//v := reflect.ValueOf(o)
	//v = v.Elem()
	//f := v.FieldByName("Age")
	//if f.Kind() == reflect.Int {
	//	fmt.Println("1111")
	//	f.SetInt(8)
	//}
	//fmt.Println("p=", p)

	//o = p
	//v := reflect.ValueOf(p)
	//m := v.MethodByName("Hello")
	//args := []reflect.Value{reflect.ValueOf("haha")}
	//m.Call(args)

	var p Person
	v := reflect.ValueOf(&p)
	t := v.Type()
	f := t.Elem().Field(0)
	fmt.Println(f.Tag.Get("json"))
	fmt.Println(f.Tag.Lookup("json"))

}

func test01() {
	var i interface{}
	num := 5
	i = &num
	t := reflect.TypeOf(i)

	fmt.Println("t=", t)
	k := t.Kind()
	fmt.Println("k=", k)

	v := reflect.ValueOf(i)
	// 指针 下这样才能打印具体值
	fmt.Println("v", v)
	fmt.Println("v.Elem()=", v.Elem())
	fmt.Println("v.Elem()=", v.Kind())
	fmt.Println("v.Elem()=", v.Type())
	fmt.Println("v.Elem()=", v.Pointer())
}
