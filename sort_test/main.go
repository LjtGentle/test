package main

import (
	"fmt"
	"sort"
	"time"
)

func test() {

}

type Student struct {
	Name string
	Age  int
}

func main() {
	s1 := Student{
		Name: "one",
		Age:  4,
	}

	s2 := Student{
		Name: "two",
		Age:  2,
	}
	s3 := Student{
		Name: "three",
		Age:  3,
	}
	s := make([]Student, 0, 3)
	s = append(s, s1)
	s = append(s, s2)
	s = append(s, s3)
	fmt.Printf("s=%+v\n", s)
	sort.Slice(s, func(i, j int) bool {
		ia := s[i].Age
		ja := s[j].Age
		return ia < ja
	})
	fmt.Printf("s=%+v\n", s)
	fmt.Println(time.Now().Unix())

}
