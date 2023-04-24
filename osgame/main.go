package main

import (
	"bufio"
	"fmt"
	"os"
	"osgame/logic"
)

type St struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//st := St{
	//	Name: "目标数量：[\"3\"]/[\"4\"]/[\"5\"]\n回复生命值：[119190p1q1]/[119191p1q1]/[119192p1q1]\\n法术伤害：[119100p1q1]/[119101p1q1]/[119102p1q1]",
	//	Age:  18,
	//}
	//b, err := json.Marshal(st)
	//if err != nil {
	//	return
	//}
	//fmt.Println("b=", string(b))
	//h := "#hhh#"
	//v := "目标数量：[\"3\"]/[\"4\"]/[\"5\"]\\n回复生命值：[119190p1q1]/[119191p1q1]/[119192p1q1]\\\\n法术伤害：[119100p1q1]/[119101p1q1]/[119102p1q1]"
	//v = strings.ReplaceAll(v, "\"", "\\\"")
	//str := "{\"name\":\"#hhh#\",\"age\":18}"
	//str = strings.ReplaceAll(str, h, v)
	//
	//err = json.Unmarshal([]byte(str), &st)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("st=%+v\n", st)
	test()
}

func test() {
	osGameTypes := [][]string{{"Standard_S1", "Basic"}}
	basic := [][]string{{"Basic", "Public"}}
	fmt.Println("osGameTypes=", osGameTypes)
	fmt.Println("basic=", basic)

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("path=", dir)
	f, err := os.Open("./ProtobufDatabin.zip")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	bin, err := logic.UpdateDataBin(f, fi.Size(), "zh_cn")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("success")
	wf, err := os.OpenFile("look.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer wf.Close()
	writer := bufio.NewWriter(wf)
	_, err = writer.WriteString(bin[0])
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("bin=", len(bin))
}
