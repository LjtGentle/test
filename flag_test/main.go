package main

import (
	"flag"
	"fmt"
	"os"
)

func test1() {
	cnfFile := flag.String("c", "./config/dev.toml", "业务配置路径")
	fmt.Printf("cnfFile =%+v\n", cnfFile)
	fmt.Printf("cnfFile =%+v\n", *cnfFile)
}

func test2() {
	var A uint64
	var B uint64
	var C uint64
	var D uint64
	flagSet := flag.FlagSet{}
	flagSet.Uint64Var(&A, "a", 1, "aaaa")
	flagSet.Uint64Var(&B, "b", 2, "bbbb")
	fmt.Println("os=", os.Args)
	flagSet.Parse(os.Args[1:])
	fmt.Printf("a=%d,b=%d,c=%d,d=%d\n", A, B, C, D)
	flagSet2 := flag.FlagSet{}
	flagSet2.Uint64Var(&C, "c", 3, "cccc")
	flagSet2.Uint64Var(&D, "d", 4, "dddd")
	flagSet2.Parse(os.Args[5:])
	fmt.Printf("a=%d,b=%d,c=%d,d=%d\n", A, B, C, D)

}

func main() {
	test2()
}
