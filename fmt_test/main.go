package main

import "fmt"

func main() {
	test()
}

func test() {
	var i interface{}
	num := 10
	i = num
	fmt.Printf("i type=%T",i)
}

