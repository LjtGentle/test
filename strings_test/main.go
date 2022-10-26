package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type People struct {
	Name string
	Age  int
}

// 装备定位
const (
	gongJi   = "1" //攻击
	faShu    = "2" //法术
	fangYu   = "3" //防御
	yiDong   = "4" //移动
	daYe     = "5" //打野
	youZou   = "7" //游走
	typeSize = 6   //6个分类
)

func test() {
	str := "aabbaabbccaaaddd"
	str2 := strings.Replace(str, "aa", "AA", -1)
	fmt.Println("str2=", str2)

}

func test2() {
	str := " 2       "
	s := strings.Split(str, "|")
	fmt.Printf("s=%+v\n", s)
}

func main() {
	test15()
}

func test15() {
	specialStr := "\\" + "\""
	fmt.Printf(specialStr)
}

//
func test14() {
	str := "hhhh\ngggg"
	str = strings.ReplaceAll(str, "\n", "")
	fmt.Println(str)
}

func test13() {
	str := "91 123 34 108 97 98 101 108 34 58 34 97 49 40 97 84 114 101 101 231 172 172 228 184 128 228 184 170 232 138 130 231 130 185 41 34 44 34 118 97 108 117 101 34 58 34 97 49 34 125 44 123 34 108 97 98 101 108 34 58 34 97 50 40 97 84 114 101 101 231 172 172 50 228 184 170 232 138 130 231 130 185 41 34 44 34 118 97 108 117 101 34 58 34 97 50 34 125 44 123 34 108 97 98 101 108 34 58 34 97 51 40 97 84 114 101 101 231 172 172 51 228 184 170 232 138 130 231 130 185 41 34 44 34 118 97 108 117 101 34 58 34 97 49 34 125 93"
	str = strings.ReplaceAll(str, " ", ",")
	fmt.Printf("str=%s\n", str)

}
func test12() {
	//var buff = []byte{123, 34, 105, 116, 101, 109, 115, 34, 58, 91, 123, 34, 108, 97, 98, 101, 108, 34, 58, 34, 97, 49, 40, 97, 84, 114, 101, 101, 231, 172, 172, 228, 184, 128, 228, 184, 170, 232, 138, 130, 231, 130, 185, 41, 34, 44, 34, 100, 101, 115, 99, 34, 58, 34, 97, 84, 114, 101, 101, 231, 172, 172, 228, 184, 128, 228, 184, 170, 232, 138, 130, 231, 130, 185, 34, 44, 34, 114, 101, 108, 97, 116, 105, 118, 101, 34, 58, 34, 97, 49, 34, 44, 34, 116, 121, 112, 101, 34, 58, 55, 44, 34, 99, 104, 105, 108, 100, 114, 101, 110, 34, 58, 123, 34, 114, 101, 102, 101, 114, 101, 110, 99, 101, 95, 97, 112, 112, 105, 100, 34, 58, 34, 116, 101, 115, 116, 58, 97, 84, 114, 101, 101, 95, 105, 100, 34, 44, 34, 114, 101, 102, 101, 114, 101, 110, 99, 101, 95, 115, 99, 104, 101, 109, 97, 105, 100, 34, 58, 34, 97, 84, 114, 101, 101, 95, 115, 99, 104, 101, 109, 97, 95, 105, 100, 34, 44, 34, 114, 101, 102, 101, 114, 101, 110, 99, 101, 95, 102, 105, 101, 108, 100, 34, 58, 34, 97, 49, 34, 44, 34, 102, 111, 114, 101, 105, 103, 110, 95, 97, 112, 112, 105, 100, 34, 58, 34, 116, 101, 115, 116, 58, 98, 84, 114, 101, 101, 95, 105, 100, 34, 44, 34, 102, 111, 114, 101, 105, 103, 110, 95, 115, 99, 104, 101, 109, 97, 105, 100, 34, 58, 34, 98, 84, 114, 101, 101, 95, 115, 99, 104, 101, 109, 97, 95, 105, 100, 34, 44, 34, 102, 111, 114, 101, 105, 103, 110, 95, 102, 105, 101, 108, 100, 34, 58, 34, 98, 50, 34, 44, 34, 109, 117, 108, 116, 105, 112, 108, 101, 34, 58, 102, 97, 108, 115, 101, 125, 125, 44, 123, 34, 108, 97, 98, 101, 108, 34, 58, 34, 97, 50, 40, 97, 84, 114, 101, 101, 231, 172, 172, 50, 228, 184, 170, 232, 138, 130, 231, 130, 185, 41, 34, 44, 34, 100, 101, 115, 99, 34, 58, 34, 97, 84, 114, 101, 101, 231, 172, 172, 50, 228, 184, 170, 232, 138, 130, 231, 130, 185, 34, 44, 34, 114, 101, 108, 97, 116, 105, 118, 101, 34, 58, 34, 97, 50, 34, 44, 34, 116, 121, 112, 101, 34, 58, 55, 44, 34, 99, 104, 105, 108, 100, 114, 101, 110, 34, 58, 110, 117, 108, 108, 125, 44, 123, 34, 108, 97, 98, 101, 108, 34, 58, 34, 97, 51, 40, 97, 84, 114, 101, 101, 231, 172, 172, 51, 228, 184, 170, 232, 138, 130, 231, 130, 185, 41, 34, 44, 34, 100, 101, 115, 99, 34, 58, 34, 97, 84, 114, 101, 101, 231, 172, 172, 51, 228, 184, 170, 232, 138, 130, 231, 130, 185, 34, 44, 34, 114, 101, 108, 97, 116, 105, 118, 101, 34, 58, 34, 97, 49, 34, 44, 34, 116, 121, 112, 101, 34, 58, 51, 44, 34, 99, 104, 105, 108, 100, 114, 101, 110, 34, 58, 110, 117, 108, 108, 125, 93, 44, 34, 110, 101, 103, 97, 116, 105, 118, 101, 34, 58, 110, 117, 108, 108, 125}
	//buff = []byte{123, 34, 105, 116, 101, 109, 115, 34, 58, 91, 123, 34, 108, 97, 98, 101, 108, 34, 58, 34, 97, 49, 40, 97, 84, 114, 101, 101, 231, 172, 172, 228, 184, 128, 228, 184, 170, 232, 138, 130, 231, 130, 185, 41, 34, 44, 34, 100, 101, 115, 99, 34, 58, 34, 97, 84, 114, 101, 101, 231, 172, 172, 228, 184, 128, 228, 184, 170, 232, 138, 130, 231, 130, 185, 34, 44, 34, 114, 101, 108, 97, 116, 105, 118, 101, 34, 58, 34, 97, 49, 34, 44, 34, 116, 121, 112, 101, 34, 58, 55, 44, 34, 99, 104, 105, 108, 100, 114, 101, 110, 34, 58, 123, 34, 114, 101, 102, 101, 114, 101, 110, 99, 101, 95, 97, 112, 112, 105, 100, 34, 58, 34, 116, 101, 115, 116, 58, 97, 84, 114, 101, 101, 95, 105, 100, 34, 44, 34, 114, 101, 102, 101, 114, 101, 110, 99, 101, 95, 115, 99, 104, 101, 109, 97, 105, 100, 34, 58, 34, 97, 84, 114, 101, 101, 95, 115, 99, 104, 101, 109, 97, 95, 105, 100, 34, 44, 34, 114, 101, 102, 101, 114, 101, 110, 99, 101, 95, 102, 105, 101, 108, 100, 34, 58, 34, 97, 49, 34, 44, 34, 102, 111, 114, 101, 105, 103, 110, 95, 97, 112, 112, 105, 100, 34, 58, 34, 116, 101, 115, 116, 58, 98, 84, 114, 101, 101, 95, 105, 100, 34, 44, 34, 102, 111, 114, 101, 105, 103, 110, 95, 115, 99, 104, 101, 109, 97, 105, 100, 34, 58, 34, 98, 84, 114, 101, 101, 95, 115, 99, 104, 101, 109, 97, 95, 105, 100, 34, 44, 34, 102, 111, 114, 101, 105, 103, 110, 95, 102, 105, 101, 108, 100, 34, 58, 34, 98, 50, 34, 44, 34, 109, 117, 108, 116, 105, 112, 108, 101, 34, 58, 102, 97, 108, 115, 101, 125, 125, 44, 123, 34, 108, 97, 98, 101, 108, 34, 58, 34, 97, 50, 40, 97, 84, 114, 101, 101, 231, 172, 172, 50, 228, 184, 170, 232, 138, 130, 231, 130, 185, 41, 34, 44, 34, 100, 101, 115, 99, 34, 58, 34, 97, 84, 114, 101, 101, 231, 172, 172, 50, 228, 184, 170, 232, 138, 130, 231, 130, 185, 34, 44, 34, 114, 101, 108, 97, 116, 105, 118, 101, 34, 58, 34, 97, 50, 34, 44, 34, 116, 121, 112, 101, 34, 58, 55, 44, 34, 99, 104, 105, 108, 100, 114, 101, 110, 34, 58, 110, 117, 108, 108, 125, 44, 123, 34, 108, 97, 98, 101, 108, 34, 58, 34, 97, 51, 40, 97, 84, 114, 101, 101, 231, 172, 172, 51, 228, 184, 170, 232, 138, 130, 231, 130, 185, 41, 34, 44, 34, 100, 101, 115, 99, 34, 58, 34, 97, 84, 114, 101, 101, 231, 172, 172, 51, 228, 184, 170, 232, 138, 130, 231, 130, 185, 34, 44, 34, 114, 101, 108, 97, 116, 105, 118, 101, 34, 58, 34, 97, 49, 34, 44, 34, 116, 121, 112, 101, 34, 58, 51, 44, 34, 99, 104, 105, 108, 100, 114, 101, 110, 34, 58, 110, 117, 108, 108, 125, 93, 44, 34, 110, 101, 103, 97, 116, 105, 118, 101, 34, 58, 110, 117, 108, 108, 125}
	var buff = []byte{91, 123, 34, 108, 97, 98, 101, 108, 34, 58, 34, 97, 49, 40, 97, 84, 114, 101, 101, 231, 172, 172, 228, 184, 128, 228, 184, 170, 232, 138, 130, 231, 130, 185, 41, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 34, 97, 49, 34, 125, 44, 123, 34, 108, 97, 98, 101, 108, 34, 58, 34, 97, 50, 40, 97, 84, 114, 101, 101, 231, 172, 172, 50, 228, 184, 170, 232, 138, 130, 231, 130, 185, 41, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 34, 97, 50, 34, 125, 44, 123, 34, 108, 97, 98, 101, 108, 34, 58, 34, 97, 51, 40, 97, 84, 114, 101, 101, 231, 172, 172, 51, 228, 184, 170, 232, 138, 130, 231, 130, 185, 41, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 34, 97, 49, 34, 125, 93}
	str := string(buff)
	fmt.Println("str=", str)
}

func test11() {
	TagIDList := make([]int, 0, 10)
	TagIDList = append(TagIDList, 1)
	TagIDList = append(TagIDList, 2)
	TagIDList = append(TagIDList, 3)
	TagIDList = append(TagIDList, 4)
	TagIDList = append(TagIDList, 5)
	fmt.Println("where", strings.Repeat("OR FIND_IN_SET(?, tags) ", len(TagIDList))[3:])
}

func test10() {
	is := make([]int, 0, 10)
	is = append(is, 10)
	is = append(is, 20)
	fmt.Printf("is=%+v\n", is)
	test9(&is)
	fmt.Printf("is=%+v\n", is)
}

func test9(is *[]int) {
	*is = append(*is, 1)
	*is = append(*is, 2)
	*is = append(*is, 3)
}

func test8() {
	ps := make([]People, 0, 10)
	p1 := People{
		Name: "tom",
		Age:  12,
	}
	p2 := People{
		Name: "jack",
		Age:  12,
	}
	p3 := People{
		Name: "ben",
		Age:  12,
	}
	ps = append(ps, p1)
	ps = append(ps, p2)
	ps = append(ps, p3)
	fmt.Printf("ps=%+v\n", ps)
	pt := make([]*People, 0, 10)
	for _, v := range ps {
		pt = append(pt, &v)
	}
	fmt.Printf("pt=%+v\n", pt)

}

// 边遍历边删除
func test7() {
	Is := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	l := len(Is)
	for i := 0; i < l; i++ {
		fmt.Println("v=", Is[i])
		if Is[i]%2 == 1 {
			fmt.Printf("v1=%d\n", Is[i])
			pre := Is[:i]
			tail := Is[i+1:]
			Is = append(pre, tail...)
			i--
			l--
		}
	}
}

func test6() {
	str := "hello"
	strs := strings.Split(str, ",")
	fmt.Printf("strs=%+v", strs)

	ss := make([]string, 0)
	ss = append(ss, "1")
	ss = append(ss, "2")
	ss = append(ss, strs...)
	fmt.Printf("ss=%+v\n", ss)
}

func test5() {
	str := "[\n                    {\n                        \"content\": \"本次大赛我们为创意官们准备了“最佳创意奖”、“优秀创意奖”、“人气创意奖”和“入围创意奖”，奖励设置如下：<br>\\n（以下奖项不累加，按获得的最高奖项发放对应奖励）\",\n                        \"id\": 1,\n                        \"pic\": \"\",\n                        \"title\": \"简介\",\n                        \"type\": \"\"\n                    },\n                    {\n                        \"content\": \"落地成为游戏内回城\",\n                        \"id\": 1627467976331,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728183817-932320.png\",\n                        \"title\": \"\",\n                        \"type\": 4\n                    },\n                    {\n                        \"content\": \"最佳创意奖官方认证证书\",\n                        \"id\": 1627468699364,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728183827-451552.png\",\n                        \"title\": \"\",\n                        \"type\": 4\n                    },\n                    {\n                        \"content\": \"Mate 30 Pro手机一台\",\n                        \"id\": 1627468708731,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728183837-521910.png\",\n                        \"title\": \"\",\n                        \"type\": 4\n                    },\n                    {\n                        \"content\": \"小小宇航员鲁班充电宝一个\",\n                        \"id\": 1627468717404,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728183845-756497.png\",\n                        \"title\": \"\",\n                        \"type\": 4\n                    },\n                    {\n                        \"content\": \"王者创意官头像框\",\n                        \"id\": 1627468725683,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728183854-469603.png\",\n                        \"title\": \"\",\n                        \"type\": 4\n                    },\n                    {\n                        \"content\": \"皮肤碎片*5\",\n                        \"id\": 1627468737443,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728183905-929818.png\",\n                        \"title\": \"\",\n                        \"type\": 4\n                    },\n                    {\n                        \"content\": \"积分夺宝抵用券*5\",\n                        \"id\": 1627468746116,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728183913-461119.png\",\n                        \"title\": \"\",\n                        \"type\": 4\n                    },\n                    {\n                        \"content\": \"共创策划基地成长值*200\",\n                        \"id\": 1627468754444,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728184031-176650.png\",\n                        \"title\": \"\",\n                        \"type\": 4\n                    },\n                    {\n                        \"content\": \"优秀创意奖官方认证证书\",\n                        \"id\": 1627468766844,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728184040-886087.png\",\n                        \"title\": \"\",\n                        \"type\": 5\n                    },\n                    {\n                        \"content\": \"黑鲨4手机一台\",\n                        \"id\": 1627468841100,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728184049-297777.png\",\n                        \"title\": \"\",\n                        \"type\": 5\n                    },\n                    {\n                        \"content\": \"小小宇航员鲁班充电宝一个\",\n                        \"id\": 1627468850643,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728184058-702040.png\",\n                        \"title\": \"\",\n                        \"type\": 5\n                    },\n                    {\n                        \"content\": \"王者创意官头像框\",\n                        \"id\": 1627468859283,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728184108-643522.png\",\n                        \"title\": \"\",\n                        \"type\": 5\n                    },\n                    {\n                        \"content\": \"皮肤碎片*5\",\n                        \"id\": 1627468869805,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728184116-684892.png\",\n                        \"title\": \"\",\n                        \"type\": 5\n                    },\n                    {\n                        \"content\": \"积分夺宝抵用券*5\",\n                        \"id\": 1627468878012,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728184126-803619.png\",\n                        \"title\": \"\",\n                        \"type\": 5\n                    },\n                    {\n                        \"content\": \"共创策划基地成长值*150\",\n                        \"id\": 1627468891484,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728184137-809591.png\",\n                        \"title\": \"\",\n                        \"type\": 5\n                    },\n                    {\n                        \"content\": \"小小宇航员鲁班充电宝一个\",\n                        \"id\": 1627468897747,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728184146-686022.png\",\n                        \"title\": \"\",\n                        \"type\": 6\n                    },\n                    {\n                        \"content\": \"王者创意官头像框\",\n                        \"id\": 1627468906899,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728184155-923250.png\",\n                        \"title\": \"\",\n                        \"type\": 6\n                    },\n                    {\n                        \"content\": \"皮肤碎片*5\",\n                        \"id\": 1627468917036,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728184205-772261.png\",\n                        \"title\": \"\",\n                        \"type\": 6\n                    },\n                    {\n                        \"content\": \"积分夺宝抵用券*5\",\n                        \"id\": 1627468926213,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728184214-591942.png\",\n                        \"title\": \"\",\n                        \"type\": 6\n                    },\n                    {\n                        \"content\": \"共创策划基地成长值*100\",\n                        \"id\": 1627468935145,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728184224-264026.png\",\n                        \"title\": \"\",\n                        \"type\": 6\n                    },\n                    {\n                        \"content\": \"皮肤碎片*5\",\n                        \"id\": 1627468945787,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728184236-692928.png\",\n                        \"title\": \"\",\n                        \"type\": 7\n                    },\n                    {\n                        \"content\": \"积分夺宝抵用券*5\",\n                        \"id\": 1627468959892,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728184253-778982.png\",\n                        \"title\": \"\",\n                        \"type\": 7\n                    },\n                    {\n                        \"content\": \"共创策划基地成长值*50\",\n                        \"id\": 1627468974004,\n                        \"pic\": \"https://sy-1254960240.cos.ap-guangzhou.myqcloud.com/smoba/ingame/images/202107/20210728184301-724352.png\",\n                        \"title\": \"\",\n                        \"type\": 7\n                    }\n                ]"
	type Award struct {
		Content string      `json:"content"`
		Id      int         `json:"id"`
		Pic     string      `json:"pic"`
		Title   string      `json:"title"`
		Type    interface{} `json:"type"`
	}
	var awards []*Award
	err := json.Unmarshal([]byte(str), &awards)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Printf("v=%+v\n", awards)
	typeMap := make(map[interface{}]bool)
	for _, a := range awards {
		i, ok := a.Type.(float64)
		findType(a.Type)
		typeMap[a.Type] = true
		fmt.Printf("i=%+v, ok=%+v\n", i, ok)
		fmt.Printf("a=%+v\n", a)
	}
	for k, v := range typeMap {
		fmt.Printf("k=%+v, v=%+v\n", k, v)
	}

}

func findType(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case uint:
		fmt.Println(x, "is uint")
	case int64:
		fmt.Println(x, "is int64")
	case uint64:
		fmt.Println(x, "is uint64")
	case uint32:
		fmt.Println(x, "is uint32")
	case int32:
		fmt.Println(x, "is int32")
	case int16:
		fmt.Println(x, "is int16")
	case uint16:
		fmt.Println(x, "is uint16")
	case int8:
		fmt.Println(x, "is int8")
	case uint8:
		fmt.Println(x, "is uint8")
	case float64:
		fmt.Println(x, "is float64")
	case float32:
		fmt.Println(x, "is float32")

	default:
		fmt.Println("非数字类型")

	}
}

func test4() {
	// 测试split开头是sep打情况
	url := "//tgl-images-1254960240.cos.ap-guangzhou.myqcloud.com/sy/smoba/cyhdy/images/20211203/1745185781_1638501163941.jpg"
	sul := strings.Split(url, "//")
	if !bytes.Equal([]byte(sul[0]), []byte("")) {
		fmt.Printf("----sul[0]=%s\n", sul[0])
		return
	}
	fmt.Printf("+++sul[0]=%s\n", sul[0])

}
func test3() {
	e := &equipmentInfoSmobaCtrl{}
	infos := make([]*EquipmentInfo, 0, 2)
	i1 := &EquipmentInfo{
		Id:      1226,
		SubType: "2",
		Icon:    "icon1",
	}
	i2 := &EquipmentInfo{
		Id:      1231,
		SubType: "2",
		Icon:    "icon2",
	}
	i3 := &EquipmentInfo{
		Id:      1232,
		SubType: "2",
		Icon:    "icon3",
	}
	i4 := &EquipmentInfo{
		Id:      1239,
		SubType: "2",
		Icon:    "icon4",
	}
	i5 := &EquipmentInfo{
		Id:      1240,
		SubType: "2",
		Icon:    "icon5",
	}
	infos = append(infos, i1)
	infos = append(infos, i2)
	infos = append(infos, i3)
	infos = append(infos, i4)
	infos = append(infos, i5)
	classifySlice, err := e.classify(infos, typeSize+2)
	if err != nil {
		return
	}
	fmt.Printf("classifySlice=%+v\n", classifySlice)
	b, err := json.Marshal(&classifySlice)
	if err != nil {
		fmt.Printf("json err=%+v\n", err)
		return
	}
	fmt.Printf("b=%s\n", string(b))
}

type equipmentInfoSmobaCtrl struct {
}
type EquipmentInfo struct {
	Id      uint64 `gorm:"column:equipment_id"` //装备id
	SubType string `gorm:"column:sub_type"`     //装备的分类 1攻击 2法术 3防御 4移动 5打野 7游走
	Icon    string `gorm:"column:icon"`         // 图标
}

type Info struct {
	Id   uint64 `json:"id"`
	Icon string `json:"icon"`
}

type ClassifyEquipAdjustResp struct {
	SubType int
	Infos   []*Info `json:"infos"`
}

func (e *equipmentInfoSmobaCtrl) classify(eInfos []*EquipmentInfo, TypeNum int) (
	[]ClassifyEquipAdjustResp, error) {
	classify := make([]ClassifyEquipAdjustResp, TypeNum)
	for _, eInfo := range eInfos {
		types := strings.Split(eInfo.SubType, "|")
		if len(types) == 1 {
			err := e.classSwitch(classify, types[0], eInfo)
			if err != nil {
				return nil, err
			}
		} else {
			for _, expr := range types {
				err := e.classSwitch(classify, expr, eInfo)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return classify, nil
}

func (e *equipmentInfoSmobaCtrl) add(class []ClassifyEquipAdjustResp, classify int, eInfo *EquipmentInfo) {
	info := &Info{
		Id:   eInfo.Id,
		Icon: eInfo.Icon,
	}
	class[classify].Infos = append(class[classify].Infos, info)
	class[classify].SubType = classify
}
func (e *equipmentInfoSmobaCtrl) classSwitch(class []ClassifyEquipAdjustResp, subType string, eInfo *EquipmentInfo) error {
	switch subType {
	case gongJi:
		e.add(class, 1, eInfo)
	case faShu:
		e.add(class, 2, eInfo)
	case fangYu:
		e.add(class, 3, eInfo)
	case yiDong:
		e.add(class, 4, eInfo)
	case daYe:
		e.add(class, 5, eInfo)
	case youZou:
		e.add(class, 7, eInfo)
	default:
		return errors.New(fmt.Sprintf("出现了未知的装备分类type=%s", subType))

	}
	return nil
}
