package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"git.code.oa.com/gap/base/mysql"
	gMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

var db *gorm.DB

func init() {
	var err error
	conStr := "root:123456@(localhost)/wuji_core?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(gMysql.Open(conStr), &gorm.Config{})
	if err != nil {
		panic(err)
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

// 无极表信息
type WujiSchema struct {
	Id               string         `gorm:"column:_id;primaryKey"`
	AppId            string         `gorm:"column:_appId;type:text" json:"_appId"`
	TableId          string         `gorm:"column:id;type:text" json:"id"`
	Name             string         `gorm:"column:name;type:text" json:"name"`
	Desc             string         `gorm:"column:desc;type:text" json:"desc"`
	Creator          string         `gorm:"column:creator;type:text" json:"creator"`
	Participant      string         `gorm:"column:participant;type:text" json:"participant"`
	SourceId         string         `gorm:"column:sourceId;type:text" json:"sourceId"`
	Private          int            `gorm:"column:private;type:tinyint(1);default:0" json:"private"`
	Actionless       int            `gorm:"column:actionless;type:tinyint(1);default:0" json:"actionless"`
	PlatformSchema   int            `gorm:"column:platformSchema;type:tinyint(1);default:0" json:"platformSchema"`
	Accessauth       string         `gorm:"column:accessauth;type:text" json:"accessauth"`
	UpdateMonitorl5  string         `gorm:"column:updateMonitorl5;type:text" json:"updateMonitorl5"`
	Fields           SchemaFieldSet `gorm:"column:fields;type:text" json:"fields"`
	Mysql            string         `gorm:"column:mysql;type:text" json:"mysql"`
	Counter          int            `gorm:"column:_counter;type:int(11);default:0" json:"_counter"`
	Ctime            time.Time      `gorm:"column:_ctime;type:datetime" json:"_ctime"`
	Mtime            time.Time      `gorm:"column:_mtime;type:datetime" json:"_mtime"`
	UserACL          string         `gorm:"column:userACL;type:text" json:"userACL"`
	TriggerAPI       string         `gorm:"column:triggerAPI;type:text" json:"triggerAPI"`
	Version          int            `gorm:"column:version;type:int(11);default:0" json:"version"`
	VersionSettings  string         `gorm:"column:versionSettings;type:text" json:"versionSettings"`
	LinkId           string         `gorm:"column:linkId;type:text" json:"linkId"`
	Triggers         string         `gorm:"column:triggers;type:text" json:"triggers"`
	Invisible        int            `gorm:"column:invisible;type:tinyint(1);default:0" json:"invisible"`
	Lock             string         `gorm:"column:_lock;type:text" json:"_lock"`
	LockExpire       int            `gorm:"column:_lock_expire;type:int(11);default:0" json:"_lock_expire"`
	Env              string         `gorm:"column:env;type:text" json:"env"`
	Forks            string         `gorm:"column:forks;type:text" json:"forks"`
	Father           string         `gorm:"column:father;type:text" json:"father"`
	Actions          string         `gorm:"column:actions;type:text" json:"actions"`
	Template         string         `gorm:"column:template;type:text" json:"template"`
	UserIsolation    int            `gorm:"column:userIsolation;type:int(11);default:0" json:"userIsolation"`
	ExtSettings      string         `gorm:"column:extSettings;type:text" json:"extSettings"`
	PrimaryKeyFields string         `gorm:"column:primaryKeyFields;type:text" json:"primaryKeyFields"`
	PrivateFields    string         `gorm:"column:privateFields;type:text" json:"privateFields"`
	SchemaKey        string         `gorm:"column:schemaKey;type:text" json:"schemaKey"`
	LastModifiedTime int64          `gorm:"column:lastModifiedTime;type:bigint(20);default:0" json:"lastModifiedTime"`
	Type             string         `gorm:"column:type;type:text" json:"type"`
	DataValidators   string         `gorm:"column:dataValidators;type:text" json:"dataValidators"`
	Group            string         `gorm:"column:group;type:text" json:"group"`
	Versioning       string         `gorm:"column:versioning;type:text" json:"versioning"`
	Permission       string         `gorm:"column:permission;type:text" json:"permission"`
	ApproachCounter  int            `gorm:"column:approachCounter;type:int(11);default:0" json:"approachCounter"`
	QpsLimit         string         `gorm:"column:qpsLimit;type:varchar(255)" json:"qpsLimit"`
	CacheSettings    string         `gorm:"column:cacheSettings;type:varchar(255)" json:"cacheSettings"`
	Relation         string         `gorm:"column:relation;type:text" json:"relation"`
}

// TableName
func (m *WujiSchema) TableName() string {
	return "_wuji_schema"
}

// SchemaField
type SchemaField struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	Desc            string        `json:"desc"`
	Type            int           `json:"type"`
	Required        bool          `json:"required"`
	Args            string        `json:"args"`
	Maxlength       interface{}   `json:"maxlength"`
	Pattern         string        `json:"pattern"`
	Sort            interface{}   `json:"sort"`
	Readonly        bool          `json:"readonly"`
	Unique          bool          `json:"unique"`
	NotAutoSpace    bool          `json:"notAutoSpace"`
	IsHidden        bool          `json:"isHidden"`
	Width           int           `json:"width"`
	Triggers        string        `json:"triggers"`
	Searchable      bool          `json:"searchable"`
	DefaultOperator []interface{} `json:"defaultOperator"`
	Score           int           `json:"score"`
	Value           string             `json:"value,omitempty"`
	Input           *SchemaFieldInput  `json:"input,omitempty"`
	Output          *SchemaFieldOutput `json:"output,omitempty"`
}

type FieldType int

type SchemaFieldInput struct {
	Type string `json:"type"`
}

type SchemaFieldOutput struct {
	Calc string `json:"calc"`
}

type SchemaFieldSet []*SchemaField

// Value
func (o SchemaFieldSet) Value() (driver.Value, error) {
	return mysql.ValueJson(o)
}

// Scan
func (o *SchemaFieldSet) Scan(input interface{}) error {
	return mysql.ScanJson(input, o)
}

// Find
//func (o SchemaFieldSet) Find(id string) *SchemaField {
//	for _, e := range o {
//		if e.ID == id {
//			return e
//		}
//	}
//	return nil
//}

func test(db *gorm.DB, typeID string) {
	var award []Award
	err := db.Table(Award{}.TableName()).Where("type_id = ?", typeID).Find(&award).Error
	if err != nil {
		fmt.Printf("db err=%+v\n", err)
		//return
	}

	//if gorm.IsRecordNotFoundError(err) {
	//	fmt.Println("11111111111111")
	//}
	if err == gorm.ErrRecordNotFound {
		fmt.Println("2222222222222222")
	}
	fmt.Printf("award=%+v\n", award)
}
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
		Bid:       "smoba1",
		TypeID:    12,
		IsCustom:  1,
		AwardType: 0,
		Name:      "ben",
	}
	a2 := Award{
		Bid:       "smoba1",
		TypeID:    12,
		IsCustom:  1,
		AwardType: 0,
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
	var a []Award
	err := db.Table(Award{}.TableName()).Where("bid = ?", "hh").Find(&a).Error
	if err != nil {
		fmt.Printf("er=%+v\n", err)
	}
	if gorm.ErrRecordNotFound == err {
		fmt.Printf("11111111111111\n")
	}
	if err == gorm.ErrRecordNotFound {
		fmt.Println("2222222222222")
	}
	fmt.Printf("err=%+v,a=%+v", err, a)

}

func main() {
	test04()
}

func test04() {
	//o := new(WujiSchema)
	//err := db.Where("_id = ?", "ingame_smoba:equipment").First(o).Error
	o := make([]*WujiSchema,0)
	err := db.Find(&o).Error
	if err != nil {
		fmt.Printf("find err=%+v\n", err)
		return
	}
	data, err := json.Marshal(o)
	if err != nil {
		fmt.Printf("json marshal err=%+v\n", err)
		return
	}
	fmt.Println("data string=", string(data))

	f, err := os.OpenFile("./schema.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("open file err=%+v", err)
	}
	defer f.Close()
	_, err = f.WriteString(string(data))
	if err != nil {
		fmt.Printf("writer err=%+v\n", err)
		return
	}
}
