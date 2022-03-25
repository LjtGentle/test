package main

import "fmt"

func test() {
	testChan := make(chan int,5)
	testChan <- 1
	testChan <- 2
	testChan <- 3
	testChan <- 4
	testChan <- 5
	l := len(testChan)
	for l > 0{
		t := <-testChan
		fmt.Println("t=",t)
		l--
	}

}

func main() {
	test()
}
