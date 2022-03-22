package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//file()
	var timer time.Time
	flag := timer.IsZero()
	fmt.Printf("flag=%+v,timer=%+v\n",flag,timer)
	stageIDs:="7,8,9,10"
	sqlStr := fmt.Sprintf("select distinct hero_id from hero_adjustment_heros where stage_id in (%s)",stageIDs)
	fmt.Println("sqlStr=",sqlStr)
	stringf()
	fmt.Println("time=",time.Now())
}


func stringf() {
	stageIDs := make([]int,0,5)
	stageIDs = append(stageIDs,7)
	stageIDs = append(stageIDs,8)
	stageIDs = append(stageIDs,9)
	stageIDs = append(stageIDs,10)
	b := []byte(fmt.Sprintf("%v",stageIDs))
	b = b[1:len(b)-1]
	str := string(b)
	str1 := strings.Replace(str," ",",",-1)
	fmt.Printf("str1=%s\n",str1)

}


func file() {

	f, err :=os.OpenFile("file.txt",os.O_RDONLY,0666)
	if err != nil {
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	data := make([]string,0,4)
	for scanner.Scan() {
		data = append(data,scanner.Text())
	}
	var res [][]string
	for _, value:= range data {
		res = append(res,strings.Split(strings.Trim(value," ")," "))
	}
	for  i:=0; i<len(res[0]);i++ {
		for j := 0; j<len(res);j++ {
			fmt.Printf("%s ",res[j][i])
		}
		fmt.Println()
	}


}