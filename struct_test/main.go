package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Age int
	Name string
}

func test() {
	p1 := People{}
	p2 := People{
		Age:  10,
		Name: "hh",
	}
	if reflect.ValueOf(p1).IsValid() {
		fmt.Println("111111111")
	}
	if reflect.ValueOf(p2).IsValid() {
		fmt.Println("222222222")
	}

	if reflect.ValueOf(People{}).IsValid() {
		fmt.Println("33333")
	}
	fmt.Println("end")
	if p1.IsEmpty() {
		fmt.Println("123")
	}
	if p2.IsEmpty() {
		fmt.Println("456")
	}
}
func (a People) IsEmpty() bool {
	return reflect.DeepEqual(a, People{})
}

func main() {
	test()
}