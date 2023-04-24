package main

import "fmt"

func main() {
	test02()
}

func test02() {
	i1 := []int{1, 2, 3}
	i2 := []int{11, 22, 33}
	copy(i1, i2)
	fmt.Println("i1=", i1)
	fmt.Println("i2=", i2)
}

func test() {
	var i interface{}
	num := 10
	i = num
	fmt.Printf("i type=%T", i)
}
