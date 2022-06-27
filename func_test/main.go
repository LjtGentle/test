package main

import "fmt"

func main() {
	test()
}
func test() {
	fmt.Println("test")
	func() {
		fmt.Println("in func test")
	}()
}