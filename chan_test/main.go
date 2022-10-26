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

type T struct {
	ch   chan string
	done chan struct{}
}

func new() *T {
	return &T{
		ch:   make(chan string, 10),
		done: make(chan struct{}),
	}
}

func (t *T) test3() {
	//time.Sleep(2 * time.Second)
	go func() {
		i := 0
		for {
			i++
			t.ch <- "hh"
			time.Sleep(1 * time.Second)
			if i > 20 {
				break
			}
		}
	}()

	go func() {
		for {
			time.Sleep(15 * time.Second)
			t.done <- struct{}{}
		}
	}()
	select {}
}

func (t *T) test() {
	defer fmt.Println("离开test")
	i := 0
	for {
		select {
		case str := <-t.ch:
			i++
			fmt.Printf("str=%s,time=%+v\n", str, time.Now())
		case <-t.done:
			fmt.Printf("get done time = %+v\n", time.Now())
			for {
				l := len(t.ch)
				fmt.Println(l)
				if l == 0 {
					break
				}
				v := <-t.ch
				fmt.Println(v)
			}
			fmt.Println("return")
			return
		default:
			//fmt.Printf("default time=%+v\n", time.Now())
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	t := new()
	go t.test3()
	t.test()
	ch := make(chan struct{})
	ch <- struct{}{}
	select {
	case <-ch:
		fmt.Println(111)
	default:
		fmt.Println(222)
	}
	fmt.Println("end of main")

}
