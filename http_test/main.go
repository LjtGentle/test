package main

import (
	"fmt"
	"net/http"
	"strings"
)

type HostEquipmentRes struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Result       `json:"result"`
}

type Result struct {
	Data     `json:"data"`
	Fixed    `json:"fixed"`
	Realtime `json:"realtime"`
	Rules    `json:"rules"`
}

type Data struct {
	EquipIdStr string `json:"equipidstr"`
	PDate      string `json:"p_date"`
}

type Fixed struct {
	HeroID string `json:"heroid"`
	RbkID  string `json:"rbk_id"`
	Secret string `json:"secret"`
}

type Realtime struct {
}

type Rules struct {
}

func main() {
	//res, err := http.Get("http://apps.datamore.tencent-cloud.net/api/pkg_rules/result?rbk_id=8663&secret=81e4851fe4df3972c3a21432bc87f0ff&heroid=127")
	//if err != nil {
	//	return
	//}
	//defer res.Body.Close()
	//fmt_test.Println("code=",res.StatusCode)
	//b, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt_test.Printf("read err=%+v\n",err)
	//	return
	//}
	//var h *HostEquipmentRes
	//err = json.Unmarshal(b, &h)
	//if err != nil {
	//	fmt_test.Printf("json err=(%+v\n)",err)
	//	return
	//}
	//fmt_test.Printf("h=%+v\n",h)

	//var h *HostEquipmentRes
	//url := "http://apps.datamore.tencent-cloud.net/api/pkg_rules/result"
	//p, err := proxy.New(proxy.Options{
	//	ConnectStr: url,
	//	BasePath:   "",
	//})
	//if err != nil {
	//	fmt.Printf("new err=%+v\n", err)
	//	return
	//}
	//_, err = p.Get("?rbk_id=8663&secret=81e4851fe4df3972c3a21432bc87f0ff&heroid=128", &h)
	//if err != nil {
	//	fmt.Printf("call get err=%+v\n", err)
	//	return
	//}
	//fmt.Printf("h=%+v\n", h)
	//p.CloseIdleConnections()

	http.HandleFunc("/test", MakeHandler)
	http.ListenAndServe("127.0.0.1:9000", nil)
	//test1()

}

func test1() {
	str := "/gp/index.html"
	ss := strings.Split(str, "/")
	fmt.Println(ss)
}

func MakeHandler(w http.ResponseWriter, r *http.Request) {
	str := ""
	bidCookie, _ := r.Cookie("test_bid")
	if bidCookie != nil {
		str = bidCookie.Value
	} else {
		fmt.Println("nil")
	}
	fmt.Println("str=", str)
}
