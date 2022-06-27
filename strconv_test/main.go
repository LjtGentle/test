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

func main() {
	test()
}
