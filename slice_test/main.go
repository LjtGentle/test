package main

import "fmt"

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
	slice_test()
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
