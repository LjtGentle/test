package main

import (
	"fmt"
	"time"
)

func test() {
	testChan := make(chan int, 5)
	testChan <- 1
	testChan <- 2
	testChan <- 3
	testChan <- 4
	testChan <- 5
	l := len(testChan)
	for l > 0 {
		t := <-testChan
		fmt.Println("t=", t)
		l--
	}

}

func test2() {
	// 测试 chan
	testChan := make(chan int, 5)

	for i := 0; i < 10; i++ {
		go func(i int) {
			var j int
			defer func() {
				testChan <- j
			}()
			time.Sleep(1 * time.Second)
			j = i
		}(i)

	}

	sum := 0
	for t := range testChan {

		fmt.Printf("t=%+v\n", t)
		sum++
		if sum == 10 {
			break
		}
	}

	fmt.Println("end..")

}

func main() {
	test2()
}
