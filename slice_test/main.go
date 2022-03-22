package main

import "fmt"

type stu struct {
	name string
	age int
}

func slice_test() {
	sc := make([]stu,8)
	test(sc)
	fmt.Printf("sc=%+v\n",sc)

}
func test(s []stu) {
	fmt.Printf("s=%+v,s-cap=%d,s-len=%d\n",s,cap(s),len(s))
	s[1].name ="tom1"
	s[2].name ="tom2"
	s[3].name ="tom3"
	s[4].name ="tom4"
	s[5].name ="tom5"
	s[7].name ="tom7"
	fmt.Printf("s=%+v\n",s)
}

func main() {
	slice_test()
}
