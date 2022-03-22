package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)
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
	str2 := strings.Replace(str,"aa","AA",-1)
	fmt.Println("str2=",str2)

}

func test2() {
	str := " 2       "
	s := strings.Split(str,"|")
	fmt.Printf("s=%+v\n",s)
}
func main() {
	test3()
}

func test3() {
	e := &equipmentInfoSmobaCtrl{}
	infos := make([]*EquipmentInfo,0,2)
	i1 := &EquipmentInfo{
		Id:      1226,
		SubType: "2",
		Icon: "icon1",
	}
	i2 := &EquipmentInfo{
		Id:      1231,
		SubType: "2",
		Icon: "icon2",
	}
	i3 := &EquipmentInfo{
		Id:      1232,
		SubType: "2",
		Icon: "icon3",
	}
	i4 := &EquipmentInfo{
		Id:      1239,
		SubType: "2",
		Icon: "icon4",
	}
	i5 := &EquipmentInfo{
		Id:      1240,
		SubType: "2",
		Icon: "icon5",
	}
	infos = append(infos,i1)
	infos = append(infos,i2)
	infos = append(infos,i3)
	infos = append(infos,i4)
	infos = append(infos,i5)
	classifySlice, err := e.classify(infos, typeSize+2)
	if err != nil {
		return
	}
	fmt.Printf("classifySlice=%+v\n",classifySlice)
	b, err := json.Marshal(&classifySlice)
	if err != nil {
		fmt.Printf("json err=%+v\n",err)
		return
	}
	fmt.Printf("b=%s\n",string(b))
}


type equipmentInfoSmobaCtrl struct {

}
type EquipmentInfo struct {
	Id      uint64 `gorm:"column:equipment_id"`       //装备id
	SubType string `gorm:"column:sub_type"` //装备的分类 1攻击 2法术 3防御 4移动 5打野 7游走
	Icon string `gorm:"column:icon"` // 图标
}


type Info struct {
	Id uint64 `json:"id"`
	Icon string `json:"icon"`
}

type ClassifyEquipAdjustResp struct {
	SubType int
	Infos []*Info `json:"infos"`
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
	class[classify].Infos = append(class[classify].Infos,info)
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