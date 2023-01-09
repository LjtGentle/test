package main

import (
	"database/sql/driver"
	"fmt"
	"git.code.oa.com/gap/base/mysql"
	_ "github.com/go-sql-driver/mysql" // mysql驱动
	"github.com/jinzhu/gorm"
)

func main() {
	test08()
}

func test04() {
	db = db.Table(Award{}.TableName()).Where("bid = ?", "smoba")
	var idList = []uint{1, 2, 3, 4, 5}
	db = db.Where("id in (?)", idList)
	err := db.Delete(&Award{}).Error
	if err != nil {
		fmt.Println("del err=", err)
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
	Id        uint      `gorm:"column:id" json:"id"`
	Bid       string    `gorm:"column:bid" json:"bid"`
	TypeID    uint      `gorm:"column:type_id" json:"type_id"`       // 期次ID
	IsCustom  int       `gorm:"column:is_custom" json:"is_custom"`   // 奖项类型=0-默认奖项,1-自定义奖项
	AwardType int       `gorm:"column:award_type" json:"award_type"` // 默认的奖项 未获奖：0，入围奖：1， 特别鼓励奖：2，人气设计奖：3， 最佳设计奖：4， 入围创意奖：5，人气创意奖：6，优秀创意奖：7，最佳创意奖：8
	Name      AwardName `gorm:"column:name" json:"name"`             // 奖项名称
}

type AwardName struct {
	MainName string `json:"main_name"`
	SideName string `json:"side_name"`
}

func (o AwardName) Value() (driver.Value, error) {
	return mysql.ValueJson(o)
}
func (o *AwardName) Scan(input interface{}) error {
	return mysql.ScanJson(input, o)
}

func (m Award) TableName() string {
	return "test.qa_award_smoba"
}

//func test06() {
//	as := make([]*Award, 0, 10)
//	as1 := &Award{Name: "name1"}
//	as2 := &Award{Name: "name2"}
//	as = append(as, as1)
//	as = append(as, as2)
//	err := db.Model(Award{}).Create(&as1).Error
//	if err != nil {
//		fmt.Println("err=", err)
//		return
//	}
//}

func test07() {
	as := make([]*Award, 0)
	query := db.Model(Award{})
	query = query.Where("bid = ?", "bid")

	values := []interface{}{0, 0}
	query = query.Where("type_id =? or award_type=?", values...)
	err := query.Find(&as).Error
	if err != nil {
		fmt.Println("err=", err)
	}
}

func test08() {
	as := make([]Award, 0)
	err := db.Model(Award{}).Find(&as).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("v1")
	fmt.Printf("%+v\n", as)
}

func test05() {
	fmt.Println(11111)
	as := make([]*Award, 0)
	err := db.Table(Award{}.TableName()).Where("id> ?", 0).
		Or("bid like ?", "bid1").Or("bid like ?", "bid2").Where("type_id=?", 0).Find(&as).Error
	if err != nil {
		fmt.Println("err=", err)
	}
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
//func test02() {
//	a := &Award{
//		Bid:       "smoba",
//		TypeID:    12,
//		IsCustom:  1,
//		AwardType: 0,
//		Name:      "GentleAName",
//	}
//	fmt.Printf("a=%+v\n", a)
//	err := db.Table(Award{}.TableName()).Create(a).Error
//	if err != nil {
//		fmt.Printf("err=%+v\n", err)
//		return
//	}
//	fmt.Printf("a=%+v\n", a)
//	as := make([]Award, 0, 10)
//	a1 := Award{
//		Id:        100,
//		Bid:       "smoba3",
//		TypeID:    12,
//		IsCustom:  1,
//		AwardType: 1,
//		Name:      "ben",
//	}
//	a2 := Award{
//		Id:        101,
//		Bid:       "smoba3",
//		TypeID:    12,
//		IsCustom:  1,
//		AwardType: 1,
//		Name:      "jack",
//	}
//	as = append(as, a1)
//	as = append(as, a2)
//	fmt.Printf("brefor as =%+v\n", as)
//	err = db.Table(Award{}.TableName()).Create(&as).Error
//	if err != nil {
//		fmt.Printf("err222=%+v\n", err)
//		return
//	}
//	fmt.Printf("as=%+v\n", as)
//}

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
