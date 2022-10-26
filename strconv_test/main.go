package main

import (
	"fmt"
	"strconv"
)

// 数字转字符串
func test() {
	i := 20
	str := strconv.Itoa(i)
	fmt.Printf("str=%+v, type= %T", str, str)
}

func test2() {
	n1 := 1
	n2 := 3
	var f float32
	f = float32(n1) / float32(n2)
	fmt.Printf("f=%3.5f\n", f)
	str := fmt.Sprintf("%3.5f", f*100) + "%"
	fmt.Printf("str=%s,type=%T", str, str)
}

func main() {
	test2()
}
