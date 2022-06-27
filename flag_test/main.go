package main

import (
	"flag"
	"fmt"
)

func test1() {
	cnfFile := flag.String("c", "./config/dev.toml", "业务配置路径")
	fmt.Printf("cnfFile =%+v\n",cnfFile)
	fmt.Printf("cnfFile =%+v\n",*cnfFile)
}


func main() {
	test1()
}
