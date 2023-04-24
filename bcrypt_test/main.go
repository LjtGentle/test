package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	test01()
}

func test01() {
	// 加密
	password := "abc"
	sword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(sword))
	// 验证
	err = bcrypt.CompareHashAndPassword(sword, []byte(password))
	if err != nil {
		fmt.Println("wrong")
		return
	}
	fmt.Println("ok")
}
