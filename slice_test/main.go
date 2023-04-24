package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	test09()

}

func test09() {
	ImageDomain := "sy-1254960240.image.myqcloud.com"
	url := "https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202104/20210401162026-698863.gif"
	urlSlice := strings.Split(url, "//")
	paths := strings.Split(urlSlice[1], "/")
	paths[0] = ImageDomain
	path := strings.Join(paths, "/")
	urlSlice[1] = path
	url = strings.Join(urlSlice, "//")
	fmt.Println(url)
}

func test08() {
	println("-----")
	println(IsNotNum("1905236957"))
}

func IsNotNum(s string) bool {
	if s == "" {
		return true
	}
	println('0')
	println('9')
	println("+++++")
	for _, c := range s {
		fmt.Println(c)
		if c < '0' || c > '9' {
			return true
		}
	}
	return false
}

type stu struct {
	name string
	age  int
}

func test07() {
	ss := make([]*stu, 0, 10)
	s1 := &stu{
		name: "11",
		age:  11,
	}
	s2 := &stu{
		name: "22",
		age:  22,
	}
	s3 := &stu{
		name: "33",
		age:  33,
	}
	s4 := &stu{
		name: "44",
		age:  44,
	}
	ss = append(ss, s1)
	ss = append(ss, s2)
	ss = append(ss, s3)
	ss = append(ss, s4)
	ii := make([]int, 0, 10)
	mm := make(map[int]*stu, 10)
	for _, v := range ss {
		ii = append(ii, v.age)
		mm[v.age] = v
		v.age = v.age * 100
	}
	for _, v := range ss {
		fmt.Println(v)
	}
	fmt.Printf("ii=%+v\n", ii)
	ms := make([]*stu, 0, 10)
	for k, v := range mm {
		fmt.Printf("k=%+v,v=%+v\n", k, v)
		ms = append(ms, v)
	}
	fmt.Printf("ms=%+v\n", ms)
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
	fmt.Println(ss[0:1])
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
