package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // mysql驱动
	"github.com/jinzhu/gorm"
)



func main() {
	test04()
}

func test04() {
 	db =  db.Table(Award{}.TableName()).Where("bid = ?", "smoba")
 	var idList = []uint{1,2,3,4,5}
 	db = db.Where("id in (?)",idList)
	err := db.Delete(&Award{}).Error
 	if err != nil {
 		fmt.Println("del err=",err)
		return
	}
	fmt.Println("del success")
}
var db *gorm.DB

func init() {
	var err error
	conStr := "root:123456@(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", conStr)
	if err != nil {
		fmt.Printf("init db err=%+v\n", err)
		panic("init db err")
	}
	db = db.Debug()
}

type Award struct {
	Id        uint   `gorm:"column:id" json:"id"`
	Bid       string `gorm:"column:bid" json:"bid"`
	TypeID    uint   `gorm:"column:type_id" json:"type_id"`       // 期次ID
	IsCustom  int    `gorm:"column:is_custom" json:"is_custom"`   // 奖项类型=0-默认奖项,1-自定义奖项
	AwardType int    `gorm:"column:award_type" json:"award_type"` // 默认的奖项 未获奖：0，入围奖：1， 特别鼓励奖：2，人气设计奖：3， 最佳设计奖：4， 入围创意奖：5，人气创意奖：6，优秀创意奖：7，最佳创意奖：8
	Name      string `gorm:"column:name" json:"name"`             // 奖项名称
}

func (m Award) TableName() string {
	return "test.qa_award_smoba"
}

func test(db *gorm.DB, typeID string) {
	var award []Award
	err := db.Table(Award{}.TableName()).Where("type_id = ?", typeID).Find(&award).Error
	if err != nil {
		fmt.Printf("db err=%+v\n", err)
		//return
	}

	if gorm.IsRecordNotFoundError(err) {
		fmt.Println("11111111111111")
	}
	if err == gorm.ErrRecordNotFound {
		fmt.Println("2222222222222222")
	}
	fmt.Printf("award=%+v\n", award)
}
// reflect: call of reflect.Value.Interface on zero Value
func test02() {
	a := &Award{
		Bid:       "smoba",
		TypeID:    12,
		IsCustom:  1,
		AwardType: 0,
		Name:      "GentleAName",
	}
	fmt.Printf("a=%+v\n", a)
	err := db.Table(Award{}.TableName()).Create(a).Error
	if err != nil {
		fmt.Printf("err=%+v\n", err)
		return
	}
	fmt.Printf("a=%+v\n", a)
	as := make([]Award, 0, 10)
	a1 := Award{
		Id: 100,
		Bid:       "smoba3",
		TypeID:    12,
		IsCustom:  1,
		AwardType: 1,
		Name:      "ben",
	}
	a2 := Award{
		Id: 101,
		Bid:       "smoba3",
		TypeID:    12,
		IsCustom:  1,
		AwardType: 1,
		Name:      "jack",
	}
	as = append(as, a1)
	as = append(as, a2)
	fmt.Printf("brefor as =%+v\n", as)
	err = db.Table(Award{}.TableName()).Create(&as).Error
	if err != nil {
		fmt.Printf("err222=%+v\n", err)
		return
	}
	fmt.Printf("as=%+v\n", as)
}

func test03() {
	var a []*Award
	err := db.Table(Award{}.TableName()).Where("bid = ?", "hh").Where("type_id = ?", 10).Find(&a).Error
	if err != nil {
		fmt.Printf("er=%+v\n", err)
	}
	if gorm.IsRecordNotFoundError(err) {
		fmt.Printf("11111111111111\n")
	}
	if err == gorm.ErrRecordNotFound {
		fmt.Println("2222222222222")
	}

}

