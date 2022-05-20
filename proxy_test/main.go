package main

import (
	"encoding/json"
	"fmt"
	proxy "git.code.oa.com/iegm-open/go-http-proxy"
	"unsafe"
)

const (
	wuJiEquipmentAdjustInfoURLTest = "https://app.ingame.qq.com/ingamego/wuji_export/smoba_weapon_adjustment_detail?scene_key=WTYQjqBqEiHteypUGbZNzhIXfW0P8X4s&page=1&per_page=10&sandbox=123&openid=53A24BB14013B23BC4403B71E0309C5A&"
	filter                         = "filter={\"equip_id\":\"%s\"}"
)

type WuJiResp struct {
	AdjustTime string `json:"adjust_time"`
	AttributeSlice AttributeSlice`json:"attributes"`
	CompoundUrlSlice `json:"compound_url"`
	Desc             string `json:"desc"`
	EquipID          `json:"equip_id"`
	EquipUrl         string `json:"equip_url"`
	Generality       string `json:"generality"`
	Introduce        string `json:"introduce"`
	Name             string `json:"name"`
	Nature           int    `json:"nature"`
	Price            int    `json:"price"`
	PrimaryDetails   `json:"primary_details"`
	StageID          int    `json:"stage_id"`
	Title            string `json:"title"`
	Version          string `json:"version"`
}



type Attributes struct {
	Attribute string `json:"attribute"`
	Desc      string `json:"desc"`
}
//type AttributeSlice []Attributes

type AttributeSlice struct {
	Attributes []Attributes `json:"attributes"`
}

type CompoundUrl struct {
	EquipmentId int `json:"equipment_id"`
}

type CompoundUrlSlice struct {
	CompoundUrl []CompoundUrl `json:"compound_url"`
}

type EquipID struct {
	Desc        string `json:"desc"`
	EquipmentId int    `json:"equipment_id"`
	HeroID      `json:"hero_id"`
	StateSlice  `json:"state"`
	SubType     string `json:"sub_type"`
	Type        int    `json:"type"`
}

type HeroID struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type StateSlice struct {
	State []State `json:"state"`
}

type State struct {
	StateID   string `json:"state"`
	StartTime int    `json:"start_time"`
	EndTime   int    `json:"end_time"`
}

type PrimaryDetails struct {
	Primary `json:"primary"`
}

type Primary struct {
	Name       string       `json:"name"`
	Price      int          `json:"price"`
	Icon       string       `json:"icon"`
	Desc       string       `json:"desc"`
	Attributes []Attributes `json:"attributes"`
}

//type WuJiRes struct {
//	AdjustTime     time.Time      `json:"adjust_time"`
//	Attributes     AttributesRes    `json:"attributes"`
//	CompoundURL    CompoundURLRes    `json:"compound_ur"`
//	Desc           string         `json:"desc"`
//	EquipID        EquipID        `json:"equip_id"`
//	EquipURL       string         `json:"equip_url"`
//	Generality     string         `json:"generality"`
//	Introduce      string         `json:"introduce"`
//	Name           string         `json:"name"`
//	Nature         int            `json:"nature"`
//	Price          int            `json:"price"`
//	PrimaryDetails PrimaryDetails `json:"primary_details"`
//	StageID        int            `json:"stage_id"`
//	Title          string         `json:"title"`
//	Version        string         `json:"version"`
//}
//type Attributes struct {
//	Attribute string `json:"attribute"`
//	Desc      string `json:"desc"`
//}
//type AttributesRes struct {
//	Attributes []Attributes `json:"attributes"`
//}
//type CompoundURL struct {
//	EquipmentID int `json:"equipment_id"`
//}
//type CompoundURLRes struct {
//	CompoundURL []CompoundURL `json:"compound_url"`
//}
//type HeroID struct {
//	ID   int    `json:"id"`
//	Name string `json:"name"`
//}
//type State struct {
//	State     string `json:"state"`
//	StartTime int    `json:"start_time"`
//	EndTime   int    `json:"end_time"`
//}
//type StateRes struct {
//	State []State `json:"state"`
//}
//type EquipID struct {
//	Desc        string `json:"desc"`
//	EquipmentID int    `json:"equipment_id"`
//	HeroID      HeroID `json:"hero_id"`
//	State       StateRes  `json:"state"`
//	SubType     string `json:"sub_type"`
//	Type        int    `json:"type"`
//}
//type Primary struct {
//	Name       string       `json:"name"`
//	Price      int          `json:"price"`
//	Icon       string       `json:"icon"`
//	Desc       string       `json:"desc"`
//	Attributes []Attributes `json:"attributes"`
//}
//type PrimaryDetails struct {
//	Primary Primary `json:"primary"`
//}

///
type Host struct {
	Date       string      `json:"date"`
	HeroID     string      `json:"hero_id"`
	Equipments []interface{} `json:"equipments"`
}

type Equipment struct {
	EquipID string `json:"equip_id"`
	Nature  int    `json:"nature"`
	Info    string `json:"info"`
}

func test() {
	wuJiProxy, err := proxy.New(proxy.Options{
		ConnectStr: wuJiEquipmentAdjustInfoURLTest,
		BasePath:   "",
	})
	if err != nil {
		fmt.Println("111111111err=", err)
		return
	}
	//infos := make([]WuJiResp, 1)
	//var infos interface{}
	//infos := make([]WuJiResp,1)
	var infos string
	code, err := wuJiProxy.Get(fmt.Sprintf(filter, "21713"), &infos)
	if err != nil {
		fmt.Printf("2222222,code=%d err=%+v\n", code, err)
		return
	}
	fmt.Println("code=", code)
	fmt.Printf("infos=%+v\n", infos)
	//infoSlice := ParseAsSerializedJSONArray(infos)
	//fmt_test.Printf("infoSlice=%+v\n",infoSlice)
	h := &Host{
		Date:       "20220317",
		HeroID:     "106",
	}
	err = json.Unmarshal([]byte(infos), &h.Equipments)
	if err != nil {
		fmt.Printf("json unmarshal err=%+v\n", err)
		return
	}

	b, err := json.Marshal(h)
	if err != nil {
		return
	}
	fmt.Printf("b=%+v\n",string(b))
	//hostProxy, err := proxy.New(proxy.Options{
	//	ConnectStr: "https://tiem-cdn.qq.com/ingame/smoba/adjustment/test_host_equipment_106.json",
	//	BasePath:   "",
	//})
	//var h Host
	//_, err = hostProxy.Get("", &h)
	//if err != nil {
	//	fmt_test.Printf("3333333333err=%+v\n", err)
	//	return
	//}
	//fmt_test.Printf("h=%+v\n", h)
}

func main() {
	test()
}

type MyAttributeSlice struct {
	AttributeSlice AttributeSlice `json:"attributes"`
}

func test2() {
	str :="[{\"adjust_time\":\"2022-01-13T00:00:00Z\",\"attributes\":{\"attributes\":[{\"attribute\":\"被动-专精(张飞）\",\"desc\":\"【守护机关】和【崩裂践踏】释放后对附近的敌人持续造成法术伤害\\u003cbr\\u003e【守护机关】可对敌人造成物理伤害，自身获得护盾增加，但不再对队友生效\"},{\"attribute\":\"被动-力量进阶\",\"desc\":\"当额外物理攻击达到200时，【守护机关】和【崩裂践踏】命中后还可减少目标150物理防御\"},{\"attribute\":\"限制条件\",\"desc\":\"失去专精装效果后5分钟内不可再次购买\"}]},\"compound_url\":{\"compound_url\":[{\"equipment_id\":0}]},\"desc\":\"1111\",\"equip_id\":{\"desc\":\"\\u003cp\\u003e通过购买专精装，张飞可以将自己切换至进攻性方向。人形状态时，他的2技能将不再为队友提供护盾，但同时，2技能落地时会对周围敌人造成物理伤害。同时，人形状态与变身状态下，释放2技能后，还会对周围敌人持续造成魔法伤害。\\u003cbr\\u003e装备专精装的状态下，张飞的保护能力会有所降低，但是进攻能力与清兵能力都有提升，为偏好进攻的玩家，或者需要主动进攻、清兵的对局，提供了选择。\\u003c/p\\u003e\",\"equipment_id\":21713,\"hero_id\":{\"id\":171,\"name\":\"张飞\"},\"state\":{\"state\":[{\"state\":\"1\",\"start_time\":1641005608,\"end_time\":1703991208}]},\"sub_type\":\"3|7\",\"type\":2},\"equip_url\":\"https://wuji-1254960240.file.myqcloud.com/xy/ingame_sy/smoba/equipment_icon/web5fbe2dae-71a3-449b-fd7f-ae40b50821bc.png\",\"generality\":\"\",\"introduce\":\"\\u003cp\\u003e通过购买专精装，张飞可以将自己切换至进攻性方向。人形状态时，他的2技能将不再为队友提供护盾，但同时，2技能落地时会对周围敌人造成物理伤害。同时，人形状态与变身状态下，释放2技能后，还会对周围敌人持续造成魔法伤害。\\u003cbr\\u003e装备专精装的状态下，张飞的保护能力会有所降低，但是进攻能力与清兵能力都有提升，为偏好进攻的玩家，或者需要主动进攻、清兵的对局，提供了选择。\\u003c/p\\u003e\",\"name\":\"震军虎啸\",\"nature\":1,\"price\":2060,\"primary_details\":{\"primary\":{\"name\":\"虎啸\",\"price\":300,\"icon\":\"https://wuji-1254960240.file.myqcloud.com/xy/ingame_sy/webe2665ab9-425a-482a-892b-cf9807d78b97.png\",\"desc\":\"\",\"attributes\":[{\"attribute\":\"被动-专精(张飞)\",\"desc\":\"【守护机关】和【崩裂践踏】释放后对附近的敌人持续造成法术伤害\\u003cbr\\u003e【守护机关】可对敌人造成物理伤害，自身获得护盾增加，但不再对队友生效\"}]}},\"stage_id\":72,\"title\":\"张飞专精装上线\",\"version\":\"25\"},{\"adjust_time\":\"2022-01-17T00:00:00Z\",\"attributes\":[{\"name\":\"\",\"desc\":\"\"}],\"compound_url\":[{\"equip_id\":0}],\"desc\":\"222\",\"equip_id\":{\"desc\":\"\\u003cp\\u003e通过购买专精装，张飞可以将自己切换至进攻性方向。人形状态时，他的2技能将不再为队友提供护盾，但同时，2技能落地时会对周围敌人造成物理伤害。同时，人形状态与变身状态下，释放2技能后，还会对周围敌人持续造成魔法伤害。\\u003cbr\\u003e装备专精装的状态下，张飞的保护能力会有所降低，但是进攻能力与清兵能力都有提升，为偏好进攻的玩家，或者需要主动进攻、清兵的对局，提供了选择。\\u003c/p\\u003e\",\"equipment_id\":21713,\"hero_id\":{\"id\":171,\"name\":\"张飞\"},\"state\":{\"state\":[{\"state\":\"1\",\"start_time\":1641005608,\"end_time\":1703991208}]},\"sub_type\":\"3|7\",\"type\":2},\"equip_url\":\"https://wuji-1254960240.file.myqcloud.com/xy/ingame_sy/smoba/equipment_icon/web5fbe2dae-71a3-449b-fd7f-ae40b50821bc.png\",\"generality\":\"\",\"introduce\":\"\\u003cp\\u003enull\\u003c/p\\u003e\",\"name\":\"震军虎啸\",\"nature\":2,\"price\":2060,\"primary_details\":{\"primary\":{\"name\":\"虎啸\",\"price\":300,\"icon\":\"https://wuji-1254960240.file.myqcloud.com/xy/ingame_sy/webe2665ab9-425a-482a-892b-cf9807d78b97.png\",\"desc\":\"\",\"attributes\":[{\"attribute\":\"\",\"desc\":\"\"}]}},\"stage_id\":74,\"title\":null,\"version\":\"25\"}]"
	//strs:= ParseAsSerializedJSONArray(str)
	//fmt_test.Printf("strs=%+v\n",strs)
	var as []WuJiResp
	err := json.Unmarshal([]byte(str), &as)
	if err != nil {
		fmt.Printf("json err=%+v\n",err)
		return
	}
	fmt.Printf("as=%+v\n",as)
}

//
func ParseAsSerializedJSONArray(s string) []string {
	b := *(*[]byte)(unsafe.Pointer(&s))
	l, r := b[0], b[len(b)-1]
	if l == 91 && r == 93 {
		var a []interface{}
		err := json.Unmarshal(b, &a)
		if err != nil {
			return []string{}
		}
		result := make([]string, 0, len(a))
		for _, item := range a {
			result = append(result, fmt.Sprintf("%v", item))
		}
	}
	return nil
}
