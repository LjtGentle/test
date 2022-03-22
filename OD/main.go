package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var num,flag,lx,lxflag int
	fmt.Scanf("%d",&num)

	var absent,late,leaveearly,present int

	jls := make([]string,num)

	scanner := bufio.NewScanner(os.Stdin)

	for i:=0;i<num;i++ {
		scanner.Scan()
		jl := scanner.Text()
		jls = strings.Split(jl," ")

		for _,v := range jls {
			switch v {
			case "absent":absent++
			case "late":late++
			case "leaveearly":leaveearly++
			case "present":present++
			default:
			}

			if v == "late" || v == "leaveearly" {
				lx++
			} else {
				lx=0
			}

			if lx > 1 {
				lxflag =1
			}

			if present > 6 {
				flag = 1
			}

			if v != "present" {
				present = 0
			}
		}


		if absent < 2 || (absent + late + leaveearly < 4 && flag ==1) || lxflag == 0 {
			fmt.Printf("true ")
		} else {
			fmt.Printf("false ")
		}
	}
}