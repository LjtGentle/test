package main

import (
	"fmt"
	"github.com/go-ego/gpy"
	"strings"
)

func main() {
	zw := "ac中国2台湾3abc"
	pySlice := gpy.Py(zw)
	fmt.Println("pySlice=", pySlice)
	pySlice = strings.ReplaceAll(pySlice, " ", "")
	fmt.Println("pySlice=", pySlice)
}

func test() {

}
