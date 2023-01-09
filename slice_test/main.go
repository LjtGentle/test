package main

import (
	"fmt"
	"time"
)

type stu struct {
	name string
	age  int
}

func slice_test() {
	sc := make([]stu, 0, 8)
	test(sc)
	fmt.Printf("sc=%+v\n", sc)

}
func test(s []stu) {
	fmt.Printf("s=%+v,s-cap=%d,s-len=%d\n", s, cap(s), len(s))
	s1 := stu{
		name: "tom",
		age:  10,
	}
	s2 := stu{
		name: "jack",
		age:  15,
	}
	s3 := stu{
		name: "ben",
		age:  20,
	}
	s = append(s, s1)
	s = append(s, s2)
	s = append(s, s3)
	fmt.Printf("berfor call test3 s=%+v\n", s)
	test3(s)
	fmt.Printf("after call test3 s=%+v\n", s)
}

func main() {
	test06()
}

func test06() {
	ss := make([]*stu, 0, 10)
	s1 := stu{
		name: "one",
		age:  1,
	}
	s2 := stu{
		name: "two",
		age:  2,
	}
	s3 := stu{
		name: "three",
		age:  3,
	}
	ss = append(ss, &s1)
	ss = append(ss, &s2)
	ss = append(ss, &s3)

	for _, s := range ss {
		go func(st *stu) {
			time.Sleep(time.Millisecond * 200)
			fmt.Printf("st=%+v\n", st)
		}(s)
	}
	select {}
}

func test05() {
	strChan := make(chan string, 100)
	fmt.Println(len(strChan))
}

func test04() {
	ss := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	ss = append(ss[:5-1], ss[5:]...)
	fmt.Println(ss)
}

func test3(students []stu) {
	for i := 0; i < len(students); i++ {
		students[i].name = "gentle"
	}
	fmt.Printf("in test 3=%+v\n", students)
}

func test1() []int {
	is := []int{1, 2, 3, 4}
	fmt.Println("in test1 is =", is)
	return is
}

func test2() {
	var is []int
	is = test1()
	fmt.Println("in test2 is=", is)
}
