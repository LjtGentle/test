package main

import (
	"fmt"
	"regexp"
	"strings"
)

func test1() {
	reg, err := regexp.Compile(`\${[\w_]+\}`)
	if err != nil {
		fmt.Println(" regexp err")
	}
	txt := "smoba_weapon_#{stage}_${version}"

	strSlice := reg.FindAllString(txt, -1)
	for _, v := range strSlice {
		txt = strings.Replace(txt,v,"V73",-1)

	}
	fmt.Printf("strSlice=%v\n", strSlice)
	fmt.Printf("txt=%v\n", txt)
	//fmt_test.Printf("strSlice[0]=%v\n",templateGetValue(strSlice[0]))
	fmt.Println("len=", len(strSlice))
}

// #{stage} -> stage  ${version -> version
func templateGetValue(str string) string {
	b := []byte(str)
	return string(b[2 : len(b)-1])
}

func main() {
	test1()
}
