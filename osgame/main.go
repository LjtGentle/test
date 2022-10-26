package main

type Translate struct {
	Records []Record `json:"Records"`
}
type Record struct {
	Key         string `json:"Key"`
	Translation string `json:"Translation"`
}

var conRecord string = "{\"Key\":\"##\",\"Translation\":\"123\"}"

func main() {
	//url := "https://tiem-cdn.qq.com/osgame/data/Standard_S1/Basic/test_ResSkillCfgInfo.zh_cn.lang.json"
	//getReader, err := http.Get(url)
	//if err != nil {
	//	fmt.Printf("http get err=%+v\n", err)
	//	return
	//}
	//defer getReader.Body.Close()
	//readByte, err := ioutil.ReadAll(getReader.Body)
	//if err != nil {
	//	fmt.Printf("ioutil.ReadAll err=%+v\n", err)
	//	return
	//}
	//var t Translate
	//err = json.Unmarshal(readByte, &t)
	//if err != nil {
	//	fmt.Printf("readbyte json unmarshal")
	//	return
	//}
	//var rec Record
	//err = json.Unmarshal([]byte(conRecord), &rec)
	//if err != nil {
	//	fmt.Printf("conRecord json unmarshal")
	//	return
	//}
	//fmt.Println(rec)
	////fmt.Printf("t=%+v\n", t)
	////
	//
	//fmt.Println("len=", len(t.Records))
	//for _, v := range t.Records {
	//	var ttt Record
	//	v.Translation = strings.ReplaceAll(v.Translation, "\n", "")
	//	v.Translation = strings.ReplaceAll(v.Translation, "\"", "")
	//	iconRecord := strings.ReplaceAll(conRecord, "##", v.Translation)
	//	err = json.Unmarshal([]byte(iconRecord), &ttt)
	//	if err != nil {
	//		fmt.Println(iconRecord)
	//		fmt.Printf("key=%s,json unmarshal err=%+v\n", v.Key, err)
	//		return
	//	}
	//}
	//fmt.Println("success")

}

func test1() {
	//
}