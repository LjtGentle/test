package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	//test_map2()
	//testSlice()
	//test5()
	//classMap := make(map[int]string)
	//add(classMap,1,"7")
	//add(classMap,1,"8")
	//add(classMap,2,"4")
	//fmt_test.Printf("classMap=%+v\n",classMap)
	//randString()
	//jsonToStrings(`{"id":632257508,"distribute":{"channel":{"chan":1,"task":16,"stage":57}}}`)
	test32()
}

func test32() {
	jsonStr := `[
    {
        "hero_id": 155,
        "time": 1650902400
    },
    {
        "hero_id": 521,
        "time": 1668009600
    },
    {
        "hero_id": 162,
        "time": 1665244800
    },
    {
        "hero_id": 534,
        "time": 1668009600
    },
    {
        "hero_id": 548,
        "time": 1668009600
    },
    {
        "hero_id": 110,
        "time": 1669132800
    },
    {
        "hero_id": 186,
        "time": 1664380800
    },
    {
        "hero_id": 163,
        "time": -62135596800
    },
    {
        "hero_id": 148,
        "time": 1668009600
    },
    {
        "hero_id": 192,
        "time": 1668009600
    },
    {
        "hero_id": 106,
        "time": 1663776000
    },
    {
        "hero_id": 112,
        "time": 1669132800
    },
    {
        "hero_id": 544,
        "time": 1669132800
    },
    {
        "hero_id": 515,
        "time": 1664380800
    },
    {
        "hero_id": 199,
        "time": 1668009600
    },
    {
        "hero_id": 111,
        "time": 1669132800
    },
    {
        "hero_id": 527,
        "time": 1664380800
    },
    {
        "hero_id": 107,
        "time": 1663776000
    },
    {
        "hero_id": 194,
        "time": 1668009600
    },
    {
        "hero_id": 184,
        "time": 1668528000
    },
    {
        "hero_id": 156,
        "time": 1669132800
    },
    {
        "hero_id": 507,
        "time": 1668009600
    },
    {
        "hero_id": 105,
        "time": 1663776000
    },
    {
        "hero_id": 195,
        "time": 1668614400
    },
    {
        "hero_id": 130,
        "time": 1668528000
    },
    {
        "hero_id": 119,
        "time": 1669132800
    },
    {
        "hero_id": 169,
        "time": 1665244800
    },
    {
        "hero_id": 149,
        "time": 1665244800
    },
    {
        "hero_id": 139,
        "time": 1668009600
    },
    {
        "hero_id": 108,
        "time": 1663776000
    },
    {
        "hero_id": 109,
        "time": 1669305600
    }
]`
	type Hero struct {
		HeroId int   `json:"hero_id"`
		Time   int64 `json:"time"`
	}
	hs := make([]*Hero, 0)
	err := json.Unmarshal([]byte(jsonStr), &hs)
	if err != nil {
		fmt.Println(err)
		return
	}
	heroMap := make(map[int]int, 0)
	for _, v := range hs {
		if nums, ok := heroMap[v.HeroId]; ok {
			nums++
			heroMap[v.HeroId] = nums
		} else {
			heroMap[v.HeroId] = 1
		}
	}
	fmt.Println(heroMap)
}

func test31() {
	ss := "1"
	i := cast.ToInt(ss)
	fmt.Println(i)
}

func test30() {
	key := "%s_%s_%s_%s"
	a := "aa"
	b := "bb"
	c := "cc"
	d := 10
	key = fmt.Sprintf(key, a, b, c, strconv.Itoa(d))
	fmt.Println("1111->key=", key)
	//key = fmt.Sprintf(key, c, d)
	//fmt.Println("222->key=", key)
}

func test29() {
	str := "[{\"battle\":{\"Name\":\"新手必备三分攻略\",\"LordID\":[0,0,0],\"EarlyHeroes\":[{\"HeroID\":1112,\"Position\":44,\"Level\":2,\"EquipIDs\":[21106,21107,21101],\"IsCenter\":true},{\"HeroID\":1072,\"Position\":15,\"Level\":2,\"EquipIDs\":[21121,22124,21110],\"IsCenter\":true},{\"HeroID\":1412,\"Position\":35,\"Level\":2},{\"HeroID\":1282,\"Position\":14,\"Level\":2},{\"HeroID\":1272,\"Position\":34,\"Level\":2}],\"FinalHeroes\":[{\"HeroID\":1232,\"Position\":15,\"Level\":2,\"EquipIDs\":[21121,22124,21110,31109],\"IsCenter\":true},{\"HeroID\":1122,\"Position\":48,\"Level\":2,\"EquipIDs\":[51101,21106,21107],\"IsCenter\":true},{\"HeroID\":5232,\"Position\":41,\"Level\":2},{\"HeroID\":1902,\"Position\":35,\"Level\":2},{\"HeroID\":1062,\"Position\":34,\"Level\":2},{\"HeroID\":1112,\"Position\":44,\"Level\":2},{\"HeroID\":1072,\"Position\":16,\"Level\":2},{\"HeroID\":1412,\"Position\":25,\"Level\":2},{\"HeroID\":1282,\"Position\":14,\"Level\":2},{\"HeroID\":1272,\"Position\":24,\"Level\":2}],\"ID\":\"107\",\"NewComer\":true,\"PreferredEquips\":[21121,22124,21110,31109,51101]},\"desc\":\"test\",\"main_job\":{\"1\":[\"1232\",\"1072\"],\"2\":[\"1282\"],\"4\":[\"5232\",\"1902\",\"1062\",\"1412\",\"1272\"],\"5\":[\"1122\",\"1112\"]},\"chess_relative\":[{\"fetterID\":\"41\",\"len\":8,\"hero_ids\":[\"1232\",\"1902\",\"1062\",\"1112\",\"1072\",\"1412\",\"1282\",\"1272\"]},{\"fetterID\":\"4\",\"len\":2,\"hero_ids\":[\"5232\",\"1902\"]},{\"fetterID\":\"6\",\"len\":2,\"hero_ids\":[\"1072\",\"1282\"]},{\"fetterID\":\"17\",\"len\":2,\"hero_ids\":[\"1122\",\"1112\"]},{\"fetterID\":\"5\",\"len\":1,\"hero_ids\":[\"1232\"]},{\"fetterID\":\"7\",\"len\":1,\"hero_ids\":[\"1412\"]},{\"fetterID\":\"12\",\"len\":1,\"hero_ids\":[\"1122\"]},{\"fetterID\":\"16\",\"len\":1,\"hero_ids\":[\"1272\"]},{\"fetterID\":\"21\",\"len\":1,\"hero_ids\":[\"1062\"]},{\"fetterID\":\"22\",\"len\":1,\"hero_ids\":[\"1282\"]},{\"fetterID\":\"35\",\"len\":1,\"hero_ids\":[\"1122\"]},{\"fetterID\":\"38\",\"len\":1,\"hero_ids\":[\"5232\"]},{\"fetterID\":\"42\",\"len\":1,\"hero_ids\":[\"1232\"]}]}]"
	var strMap []map[string]interface{}
	err := json.Unmarshal([]byte(str), &strMap)
	if err != nil {
		fmt.Println("json unmarshal err=", err)
		return
	}
	fmt.Printf("strMap[battle]=%+v\n", strMap[0]["battle"])
	fmt.Printf("strMap[battle] type =%T\n", strMap[0]["battle"])
}

func test28() {
	l := lengthOfLongestSubstring("aab")
	fmt.Println(l)
}

func lengthOfLongestSubstring(s string) int {
	sub, tempStr := "", ""
	// key runeValue value index
	strMap := make(map[byte]int)
	for right := 0; right < len(s); right++ {
		if index, ok := strMap[s[right]]; ok {
			//出现了重复
			if len(tempStr) > len(sub) {
				sub = tempStr
			}
			strMap = make(map[byte]int)
			right = index
			tempStr = ""
			continue
		}
		strMap[s[right]] = right
		tempStr += string(s[right])
	}
	if len(tempStr) > len(sub) {
		sub = tempStr
	}
	return len(sub)
}

func test27() {
	input := make(map[int]int)
	fmt.Println("11111")
	v, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(v))
}

func test26() {
	testMap := make(map[string]string)
	testMap["a"] = "A"
	fmt.Printf("----%s----%d--\n", testMap["b"], len(testMap["b"]))
}

func test25() {
	// 测试struct 字段比string少
	type Man struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	ss := "{\"name\":\"ben\",\"age\":10,\"addr\":\"sz\"}"

	var m Man
	err := json.Unmarshal([]byte(ss), &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("m=", m)
}

func test_map() {
	tmap := make(map[string]interface{})
	str := "{\"id\":\"13\"}"
	err := json.Unmarshal([]byte(str), &tmap)
	if err != nil {
		fmt.Printf("err=%+v\n", err)
		return
	}
	fmt.Printf("tmap=%+v\n", tmap)
	v, ok := tmap["id"]
	if !ok {
		fmt.Printf("11111\n")
		return
	}
	fmt.Printf("v=%+v\n", v)
}

func test_map2() {
	classMap := make(map[int][]int64)
	classMap[1] = []int64{1, 2, 3}
	classMap[2] = []int64{1, 9, 3}
	classMap[3] = []int64{2, 9, 3}
	b, err := json.Marshal(&classMap)
	if err != nil {
		return
	}
	fmt.Printf("b=%s\n", string(b))
}

func add(classMap map[int]string, classify int, subType string) {
	v, ok := classMap[classify]
	if !ok {
		classMap[classify] = subType
	} else {
		classMap[classify] = v + "," + subType
	}
}

func testSlice() {
	ids := [5]int{1, 2, 3, 4, 5}
	is := ids[:]
	b, err := json.Marshal(is)
	if err != nil {

		return
	}
	fmt.Printf("b=%s\n", string(b))
}

//type PageInfo struct {
//	Filter   uint32 `json:"filter"`
//	Page     uint32 `json:"page"`
//	PageSize uint32 `json:"page_size"`
//	Offset   uint32 `json:"-"`
//}

//
//func test11() {
//	str := `{\"filter\":1,\"page\":2,\"page_size\":10}`
//	var p PageInfo
//	if err := json.Unmarshal([]byte(str), &p); err != nil {
//		fmt.Printf("错误的分页数据:%+v", err)
//		return
//	}
//
//}

func test24() {
	fileName := "lordSkill/LordSkill_ChangE.png"
	fileName = strings.Join(strings.Split(fileName, "/")[1:], "/")
	ts := strings.Split(fileName, "/")[1:]
	fmt.Println(ts)
	fmt.Println(fileName)
}

type APP struct {
	AppBusiness map[string]interface{} `json:"app_business"`
}

func test23() {
	str := "{\"app_business\":{\"comment\":{},\"message\":{},\"operation-new\":{},\"qq2openid\":{},\"role\":{},\"sluglogin\":{},\"userprofile\":{},\"vote3\":{}}}"
	app := APP{}
	err := json.Unmarshal([]byte(str), &app)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("app=%+v\n", app)
	v, ok := app.AppBusiness["comment"]
	fmt.Println("v=", v)

	fmt.Println("ok=", ok)
	if v == nil {
		fmt.Println("1111")
	}
	vm, ok := v.(map[string]interface{})
	fmt.Printf("ok=%+v,len(vm)=%d\n", ok, len(vm))
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)
	fmt.Printf("rv.IsNil()=%+v\trv.IsValid()=%+v\n", rv.IsNil(), rv.IsValid())
	fmt.Printf("rv=%+v\n", rv)
	fmt.Printf("rt=%+v\n", rt)
	if rv.IsZero() {
		fmt.Println("is zero")
	} else {
		fmt.Println("not zero")
	}

	if IsZeroOfUnderlyingType(v) {
		fmt.Println("zero")
	} else {
		fmt.Println("npt zero")
	}
}

func IsZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

type KV struct {
	K string `json:"k"`
	V string `json:"v"`
}

func test22() {
	value := "黄忠架起炮台，可以对全图任意位置的敌人进行范围炮击，最多炮击7次\n4208 被动1：每次架起炮台会为自己增加双防\n4209 被动2：无论是普攻还是炮台攻击，黄忠都能为自己增加一定的攻击力和暴击率，炮台期间该效果翻倍\n4210 被动3：黄忠的炮台攻击为范围攻击，周遭的敌人会遭受到正被炮台攻击的敌人伤害的1/2"
	key := "#F21ABD0C#"
	kv := KV{
		K: key,
		V: value,
	}
	data, err := json.Marshal(&kv)
	if err != nil {
		fmt.Println("111111")
		return
	}
	fmt.Printf("data=%s\n", string(data))
	str := string(data)
	str = strings.ReplaceAll(str, "\n", "\\n")
	kk := KV{}
	err = json.Unmarshal([]byte(str), &kk)
	if err != nil {
		fmt.Println("2222err=", err)
		return
	}

	fmt.Printf("kk=%+v\n", kk)

}

func test21() {
	str := "{\"sop\":{\"app_business\":{\"comment\":{},\"role\":{},\"message\":{},\"operation-new\":{},\"vote3\":{},\"sluglogin\":{},\"qq2openid\":{},\"userprofile\":{}}}}"
	jobMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), jobMap)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Printf("jobMap=%+v\n", jobMap)
}

// 结构体对象为nil转json
func test20() {
	p := new(PageInfo)

	marshal, err := json.Marshal(nil)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("marshal=", string(marshal))
	err = json.Unmarshal(marshal, p)
	if err != nil {
		fmt.Println("err2=", err)
		return
	}
	fmt.Println("p=", p)
	if p == nil {
		fmt.Println("111111")
	}
}

func jsonToStrings(jsonStr string) {
	var obj interface{}
	json.Unmarshal([]byte(jsonStr), &obj)
	fmt.Printf("obj=%+v\n", obj)
}

func test10() {
	str := strings.ToLower("index.html")
	fmt.Println("str=", str)
}

// 随机字符串
func randString() {
	for i := 0; i < 10; i++ {
		i := rand.Int31()
		fmt.Println("i=", i)
	}
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type Stu struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type Student struct {
	Name    string `json:"name"`
	Age     string `json:"age"`
	Address string `json:"address,omitempty"`
}

// test5 测试 小字段的结构体转大字段的结构体
func test5() {
	minStu := Stu{
		Name: "Gentle",
		Age:  "24",
	}
	data, err := json.Marshal(minStu)
	if err != nil {
		fmt.Println("1111111111err=", err)
		return
	}
	var maxStu Student
	err = json.Unmarshal(data, &maxStu)
	if err != nil {
		fmt.Println("2222222222222err=", err)
		return
	}
	fmt.Printf("maxStu=%+v\n", maxStu)
}

type PageInfo struct {
	Filter   uint32 `json:"filter"`
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"page_size"`
	Offset   uint32 `json:"-"`
}

func test4() {
	var p PageInfo
	str := "{\"filter\":1,\"page\":2,\"page_size\":20}"
	if err := json.Unmarshal([]byte(str), &p); err != nil {
		fmt.Printf("错误的分页数据:%v\n", str)
		return
	}
	fmt.Printf("p=%+v", p)
}

// 删除某个下标
func test3() {
	is := []int{1, 2, 3, 4, 5, 6, 7}
	intSlice := is[:]
	fmt.Println("intSlice=", intSlice)
	for k, v := range intSlice {
		fmt.Printf("k=%d,v=%d\n", k, v)
	}
	// 删除下标为3的
	index := 3
	pre := intSlice[:index]
	tail := intSlice[index+1:]
	res := make([]int, 0, 10)
	res = append(pre, tail...)
	fmt.Println("res=", res)

}

// json 转对象

func test2() {
	//str := "{\"primary\":{\"name\":\"重拳\",\"price\":452,\"icon\":\"https://wuji-1254960240.file.myqcloud.com/xy/ingame_sy/web4ea55d73-3374-4a63-bcc4-0aae6c0bf31c.png\",\"desc\":\"\",\"attributes\":[{\"attribute\":\"被动-专精（张飞）效果调整为：\",\"desc\":\"【守护机关】可以对目标区域敌人造成伤害，并附带50%减速持续1s，但无法再为队友提供护盾，【崩裂践踏】可额外对目标造成50%减速，持续1s\"}]}}"
	str := "[{\\\"name\\\":\\\"被动-专精（张飞）效果调整为：\\\",\\\"desc\\\":\\\"【守护机关】可以对目标区域敌人造成伤害，并附带50%减速持续1s，但无法再为队友提供护盾，【崩裂践踏】可额外对目标造成50%减速，持续1s\\\"},{\\\"name\\\":\\\"被动-力量进阶效果调整为：\\\",\\\"desc\\\":\\\"额外物理攻击达到100时，释放【守护机关】和【崩裂践踏】后，会持续对周围造成法术灼烧伤害，并降低受到伤害的目标30%的输出\\\"}]"
	str = strings.Replace(str, "\\", "", -1)
	fmt.Println("str=", str)
	var obj interface{}
	err := json.Unmarshal([]byte(str), &obj)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("obj=", obj)
	b, err := json.Marshal(obj)
	if err != nil {
		return
	}
	fmt.Println("b=", string(b))

}
