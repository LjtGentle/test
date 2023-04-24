package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func main() {
	test08()
}

//func test08() {
//	validate = validator.New()
//	validate.ValidateMap()
//}

// 自定义tag 对字段的参数进行数据源校验
// todo 需要中间处理 没有办法直接使用json-schema
func test08() {
	validate = validator.New()
	data := `{
    "name": "张三",
    "age": 18,
    "locked": true,
    "tags": [
        "foor",
        "bar"
    ],
    "game": "smoba",
    "qq_appid": "123",
    "wx_appid": "abc"
}`
	rule := `{
   "name":"required,min=2,max=4",
	"age":"required,min=0,max=150",
	"locked":"required",
	"tags":"dive,oneof=foor bar",
	"game":"required,min=0,max=10,my_tag1=123456",
	"qq_appid":"required,number,my_tag2",
	"wx_appid":"required,alpha,my_tag3"
}`
	dataMap := make(map[string]interface{})
	ruleMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &dataMap)
	if err != nil {
		fmt.Println("json err=", err)
		return
	}
	err = json.Unmarshal([]byte(rule), &ruleMap)
	if err != nil {
		fmt.Println("json err=", err)
		return
	}

	// do some register
	//validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
	//	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	//	if name == "-" {
	//		return ""
	//	}
	//	return name
	//})
	// 注册game

	err = validate.RegisterValidation("my_tag1", func(fl validator.FieldLevel) bool {
		value := Gamer(fl.Field().Interface().(string))
		fmt.Println("game fl.Param()=", fl.Param())
		fmt.Println("game ttt=", fl.Top())
		// panic
		//v, k, ok1, ok2 := fl.GetStructFieldOK2()
		//fmt.Printf("game GetStructFieldOK2 v:%+v,k:%+v,ok1:%+v,ok2:%+v\n", v, k, ok1, ok2)
		//
		//v, k, ok1, ok2 = fl.GetStructFieldOKAdvanced2(fl.Parent(), "rule.age")
		//fmt.Printf("game GetStructFieldOKAdvanced2 v:%+v,k:%+v,ok1:%+v,ok2:%+v\n", v, k, ok1, ok2)
		return value.IsContain()
	})
	if err != nil {
		fmt.Println("call RegisterValidation err1=", err)
		return
	}
	err = validate.RegisterValidation("my_tag2", func(fl validator.FieldLevel) bool {
		value := QQAppId(fl.Field().Interface().(string))
		return value.IsContain()
	})
	if err != nil {
		fmt.Println("call RegisterValidation err2=", err)
		return
	}
	err = validate.RegisterValidation("my_tag3", func(fl validator.FieldLevel) bool {
		value := WxAppid(fl.Field().Interface().(string))
		return value.IsContain()
	})
	if err != nil {
		fmt.Println("call RegisterValidation err3=", err)
		return
	}
	//fmt.Printf("dataMap=%+v\n", dataMap)
	//fmt.Printf("ruleMap=%+v\n", ruleMap)
	validatorMap := validate.ValidateMap(dataMap, ruleMap)
	if len(validatorMap) != 0 {
		fmt.Println("validatorMap=", validatorMap)
		return
	}
	fmt.Println("nice!")

}

type Gamer string

func (g Gamer) IsContain() bool {
	gs := []Gamer{"smoba", "qsm", "codm"}
	for _, item := range gs {
		if item == g {
			return true
		}
	}
	return false
}

type QQAppId string

func (g QQAppId) IsContain() bool {
	gs := []QQAppId{"123", "321", "1234567"}
	for _, item := range gs {
		if item == g {
			return true
		}
	}
	return false
}

type WxAppid string

func (g WxAppid) IsContain() bool {
	gs := []WxAppid{"abc", "cba", "abcdefg"}
	for _, item := range gs {
		if item == g {
			return true
		}
	}
	return false
}

// 自定义tag
func test07() {
	validate = validator.New()

	// register function to get tag name from json tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// register validation for 'User'
	// NOTE: only have to register a non-pointer type for 'User', validator
	// internally dereferences during it's type checks.
	//validate.RegisterStructValidation(UserStructLevelValidation, User{})

	// register a custom validation for user genre on a line
	// validates that an enum is within the interval
	err := validate.RegisterValidation("gender_custom_validation", func(fl validator.FieldLevel) bool {
		value := fl.Field().Interface().(Gender)
		return value.String() != "unknown"
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// build 'User' info, normally posted data etc...
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
		City:   "Unknown",
	}

	user := &User{
		FirstName:      "",
		LastName:       "",
		Age:            45,
		Email:          "Badger.Smith@gmail",
		FavouriteColor: "#000",
		Addresses:      []*Address{address},
	}

	// returns InvalidValidationError for bad validation input, nil or ValidationErrors ( []FieldError )
	err = validate.Struct(user)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			e := validationError{
				Namespace:       err.Namespace(),
				Field:           err.Field(),
				StructNamespace: err.StructNamespace(),
				StructField:     err.StructField(),
				Tag:             err.Tag(),
				ActualTag:       err.ActualTag(),
				Kind:            fmt.Sprintf("%v", err.Kind()),
				Type:            fmt.Sprintf("%v", err.Type()),
				Value:           fmt.Sprintf("%v", err.Value()),
				Param:           err.Param(),
				Message:         err.Error(),
			}

			indent, err := json.MarshalIndent(e, "", "  ")
			if err != nil {
				fmt.Println(err)
				panic(err)
			}

			fmt.Println(string(indent))
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}

	// save user to database

}
func UserStructLevelValidation(sl validator.StructLevel) {

	user := sl.Current().Interface().(User)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "fname", "FirstName", "fnameorlname", "")
		sl.ReportError(user.LastName, "lname", "LastName", "fnameorlname", "")
	}

	// plus can do more, even with different tag than "fnameorlname"
}

type validationError struct {
	Namespace       string `json:"namespace"` // can differ when a custom TagNameFunc is registered or
	Field           string `json:"field"`     // by passing alt name to ReportError like below
	StructNamespace string `json:"structNamespace"`
	StructField     string `json:"structField"`
	Tag             string `json:"tag"`
	ActualTag       string `json:"actualTag"`
	Kind            string `json:"kind"`
	Type            string `json:"type"`
	Value           string `json:"value"`
	Param           string `json:"param"`
	Message         string `json:"message"`
}

type Gender uint

const (
	Male Gender = iota + 1
	Female
	Intersex
)

func (gender Gender) String() string {
	terms := []string{"Male", "Female", "Intersex"}
	if gender < Male || gender > Intersex {
		return "unknown"
	}
	return terms[gender]
}

// User contains user information
type User struct {
	FirstName      string     `json:"fname"`
	LastName       string     `json:"lname"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `json:"e-mail" validate:"required,email"`
	FavouriteColor string     `validate:"hexcolor|rgb|rgba"`
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
	Gender         Gender     `json:"gender" validate:"required,gender_custom_validation"`
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

//// DbBackedUser User struct
//type DbBackedUser struct {
//	Name sql.NullString `validate:"required"`
//	Age  sql.NullInt64  `validate:"required"`
//}
//
//// 自定义字段类型
//func test06() {
//	validate = validator.New()
//
//	// register all sql.Null* types to use the ValidateValuer CustomTypeFunc
//	validate.RegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})
//
//	// build object for validation
//	x := DbBackedUser{Name: sql.NullString{String: "", Valid: true}, Age: sql.NullInt64{Int64: 0, Valid: false}}
//
//	err := validate.Struct(x)
//
//	if err != nil {
//		fmt.Printf("Err(s):\n%+v\n", err)
//	}
//}
//
//// ValidateValuer implements validator.CustomTypeFunc
//func ValidateValuer(field reflect.Value) interface{} {
//
//	if valuer, ok := field.Interface().(driver.Valuer); ok {
//
//		val, err := valuer.Value()
//		if err == nil {
//			return val
//		}
//		// handle the error how you want
//	}
//
//	return nil
//}
//
//// 自定义错误类型
//func test05() {
//
//	// NOTE: ommitting allot of error checking for brevity
//
//	en := en.New()
//	uni = ut.New(en, en)
//
//	// this is usually know or extracted from http 'Accept-Language' header
//	// also see uni.FindTranslator(...)
//	trans, _ := uni.GetTranslator("en")
//
//	validate = validator.New()
//	en_translations.RegisterDefaultTranslations(validate, trans)
//
//	translateAll(trans)
//	translateIndividual(trans)
//	translateOverride(trans) // yep you can specify your own in whatever locale you want!
//}
//
//// User contains user information
//type User struct {
//	FirstName      string     `validate:"required"`
//	LastName       string     `validate:"required"`
//	Age            uint8      `validate:"gte=0,lte=130"`
//	Email          string     `validate:"required,email"`
//	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
//	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
//}
//
//// Address houses a users address information
//type Address struct {
//	Street string `validate:"required"`
//	City   string `validate:"required"`
//	Planet string `validate:"required"`
//	Phone  string `validate:"required"`
//}
//
//// use a single instance , it caches struct info
//var (
//	uni      *ut.UniversalTranslator
//	validate *validator.Validate
//)
//
//func translateAll(trans ut.Translator) {
//
//	type User struct {
//		Username string `validate:"required"`
//		Tagline  string `validate:"required,lt=10"`
//		Tagline2 string `validate:"required,gt=1"`
//	}
//
//	user := User{
//		Username: "Joeybloggs",
//		Tagline:  "This tagline is way too long.",
//		Tagline2: "1",
//	}
//
//	err := validate.Struct(user)
//	if err != nil {
//
//		// translate all error at once
//		errs := err.(validator.ValidationErrors)
//
//		// returns a map with key = namespace & value = translated error
//		// NOTICE: 2 errors are returned and you'll see something surprising
//		// translations are i18n aware!!!!
//		// eg. '10 characters' vs '1 character'
//		fmt.Println(errs.Translate(trans))
//	}
//}
//
//func translateIndividual(trans ut.Translator) {
//
//	type User struct {
//		Username string `validate:"required"`
//	}
//
//	var user User
//
//	err := validate.Struct(user)
//	if err != nil {
//
//		errs := err.(validator.ValidationErrors)
//
//		for _, e := range errs {
//			// can translate each error one at a time.
//			fmt.Println(e.Translate(trans))
//		}
//	}
//}
//
//func translateOverride(trans ut.Translator) {
//
//	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
//		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
//	}, func(ut ut.Translator, fe validator.FieldError) string {
//		t, _ := ut.T("required", fe.Field())
//
//		return t
//	})
//
//	type User struct {
//		Username string `validate:"required"`
//	}
//
//	var user User
//
//	err := validate.Struct(user)
//	if err != nil {
//
//		errs := err.(validator.ValidationErrors)
//
//		for _, e := range errs {
//			// can translate each error one at a time.
//			fmt.Println(e.Translate(trans))
//		}
//	}
//}
//
//func test04() {
//	//1679452762000
//	//1679455543409997
//	//1679455581965 UnixMilli
//	//1679544988
//	tt := 1679545202718
//	t := time.UnixMilli(int64(tt))
//	fmt.Printf("t=%+v\n", t)
//	fmt.Printf("time=%+v\n", time.Now().UnixMilli())
//}
//
//func test03() {
//	rule := `{
//  "$schema": "http://json-schema.org/draft-04/schema#",
//  "type": "object",
//  "properties": {
//    "LineupGameType": {
//      "type": "number",
//      "minimum": 0,
//      "maximum": 1
//    },
//    "Replace": {
//      "type": "object",
//      "properties": {
//        "LordID": {
//          "type": "array",
//          "items": {
//            "type": "number",
//            "minimum": 1
//          }
//        },
//        "Name": {
//          "type": "string"
//        },
//        "FinalHeroes": {
//          "type": "array",
//          "items": {
//            "type": "object",
//            "properties": {
//              "HeroID": {
//                "type": "number",
//                "minimum": 1
//              },
//              "EquipIDs": {
//                "type": "array",
//                "items": {
//                  "type": "number",
//                  "minimum": 1
//                }
//              },
//              "Position": {
//                "type": "number",
//                "minimum": 11,
//                "maximum": 48
//              },
//              "Level": {
//                "type": "number",
//                "minimum": 1,
//                "maximum": 3
//              },
//              "IsCenter": {
//                "type": "string",
//                "enum": [
//                  "true",
//                  "false"
//                ]
//              }
//            },
//            "required": [
//              "HeroID",
//              "Position",
//              "Level"
//            ]
//          }
//        },
//        "PreferredEquips": {
//          "type": "array",
//          "items": {
//            "type": "number",
//            "minimum": 1
//          }
//        },
//        "ID": {
//          "type": "string"
//        },
//        "NewComer": {
//          "type": "string",
//          "enum": [
//            "true",
//            "false"
//          ]
//        },
//        "Advance": {
//          "type": "string",
//          "enum": [
//            "true",
//            "false"
//          ]
//        },
//        "EarlyHeroes": {
//          "type": "array",
//          "items": {
//            "type": "object",
//            "properties": {
//              "HeroID": {
//                "type": "number",
//                "minimum": 1
//              },
//              "Position": {
//                "type": [
//                  "number"
//                ],
//                "minimum": 11,
//                "maximum": 48
//              },
//              "Level": {
//                "type": "number",
//                "minimum": 1,
//                "maximum": 3
//              },
//              "EquipIDs": {
//                "type": "array",
//                "items": {
//                  "type": "number",
//                  "minimum": 1
//                }
//              }
//            },
//            "required": [
//              "HeroID",
//              "Position",
//              "Level"
//            ]
//          }
//        }
//      }
//    },
//    "Index": {
//      "type": "number"
//    }
//  }
//}`
//	str := `{
//    "LineupGameType": 0,
//    "Replace": {
//        "LordID": [
//            11,
//            22,
//            32
//        ],
//        "Name": "尧天刺",
//        "FinalHeroes": [
//            {
//                "HeroID": 1672,
//                "EquipIDs": [
//                    51113,
//                    21116,
//                    21101
//                ],
//                "Position": 46,
//                "Level": 2,
//                "IsCenter": "true"
//            },
//            {
//                "HeroID": 5072,
//                "EquipIDs": [
//                    21107,
//                    21101,
//                    22118
//                ],
//                "Position": 41,
//                "Level": 2,
//                "IsCenter": "true"
//            },
//            {
//                "HeroID": 1462,
//                "EquipIDs": [
//                    21103,
//                    21108,
//                    21113,
//                    31111
//                ],
//                "Position": 43,
//                "Level": 2
//            },
//            {
//                "HeroID": 1761,
//                "Position": 36,
//                "Level": 1
//            },
//            {
//                "HeroID": 1542,
//                "EquipIDs": [
//                    21121,
//                    22124,
//                    31103
//                ],
//                "Position": 14,
//                "Level": 2
//            },
//            {
//                "HeroID": 1312,
//                "Position": 45,
//                "Level": 2
//            },
//            {
//                "HeroID": 5022,
//                "Position": 44,
//                "Level": 2
//            },
//            {
//                "HeroID": 5112,
//                "Position": 17,
//                "Level": 2
//            },
//            {
//                "HeroID": 5012,
//                "Position": 23,
//                "Level": 2
//            }
//        ],
//        "PreferredEquips": [
//            21121,
//            22124,
//            31103,
//            51113,
//            21116
//        ],
//        "ID": "90",
//        "NewComer": "true",
//        "Advance": "false",
//        "EarlyHeroes": [
//            {
//                "HeroID": 5052,
//                "Position": 15,
//                "Level": 2
//            },
//            {
//                "HeroID": 5112,
//                "EquipIDs": [
//                    21121,
//                    22124,
//                    31103
//                ],
//                "Position": 16,
//                "Level": 2
//            },
//            {
//                "HeroID": 1312,
//                "EquipIDs": [
//                    21116,
//                    21101,
//                    21107
//                ],
//                "Position": 43,
//                "Level": 2
//            },
//            {
//                "HeroID": 5022,
//                "Position": "http://tst.qq.com\\t.appt.ingame.qq.com/ssrfauto/08c4fdd0f436fcd670803f342afdb9b3",
//                "Level": 2
//            },
//            {
//                "HeroID": 1952,
//                "Position": 46,
//                "Level": 2
//            },
//            {
//                "HeroID": 1412,
//                "EquipIDs": [
//                    21103
//                ],
//                "Position": 44,
//                "Level": 2
//            },
//            {
//                "HeroID": 5012,
//                "Position": 26,
//                "Level": 2
//            }
//        ]
//    },
//    "Index": 9
//}`
//	ruleLoader := gojsonschema.NewStringLoader(rule)
//	strLoader := gojsonschema.NewStringLoader(str)
//	result, err := gojsonschema.Validate(ruleLoader, strLoader)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	if result.Valid() {
//		fmt.Printf("The document is valid\n")
//	} else {
//		fmt.Printf("The document is not valid. see errors :\n")
//		for _, err := range result.Errors() {
//			// Err implements the ResultError interface
//			fmt.Printf("- %s\n", err)
//		}
//	}
//}
//
//func test02() {
//
//	v := validator.New()
//	ss := "{\"data\":[{\"token\":\"88\"}]}"
//	ssMap := make(map[string]interface{})
//	err := json.Unmarshal([]byte(ss), &ssMap)
//	if err != nil {
//		return
//	}
//	rules := map[string]interface{}{
//		"data": map[string]interface{}{
//			"token": "required,min=0",
//		},
//	}
//
//	m := v.ValidateMap(ssMap, rules)
//	if m != nil {
//		fmt.Printf("%+v\n", m)
//	}
//
//}
//
//func test01() {
//	str := "{\"LineupGameType\":0,\"Replace\":{\"LordID\":[11,22,32],\"Name\":\"尧天刺\",\"FinalHeroes\":[{\"HeroID\":1672,\"EquipIDs\":[51113,21116,21101],\"Position\":46,\"Level\":2,\"IsCenter\":\"true\"},{\"HeroID\":5072,\"EquipIDs\":[21107,21101,22118],\"Position\":41,\"Level\":2,\"IsCenter\":\"true\"},{\"HeroID\":1462,\"EquipIDs\":[21103,21108,21113,31111],\"Position\":43,\"Level\":2},{\"HeroID\":1761,\"Position\":36,\"Level\":1},{\"HeroID\":1542,\"EquipIDs\":[21121,22124,31103],\"Position\":14,\"Level\":2},{\"HeroID\":1312,\"Position\":45,\"Level\":2},{\"HeroID\":5022,\"Position\":44,\"Level\":2},{\"HeroID\":5112,\"Position\":17,\"Level\":2},{\"HeroID\":5012,\"Position\":23,\"Level\":2}],\"PreferredEquips\":[21121,22124,31103,51113,21116],\"ID\":\"90\",\"NewComer\":\"true\",\"Advance\":\"false\",\"EarlyHeroes\":[{\"HeroID\":5052,\"Position\":15,\"Level\":2},{\"HeroID\":5112,\"EquipIDs\":[21121,22124,31103],\"Position\":16,\"Level\":2},{\"HeroID\":1312,\"EquipIDs\":[21116,21101,21107],\"Position\":43,\"Level\":2},{\"HeroID\":5022,\"Position\":40,\"Level\":2},{\"HeroID\":1952,\"Position\":46,\"Level\":2},{\"HeroID\":1412,\"EquipIDs\":[21103],\"Position\":44,\"Level\":2},{\"HeroID\":5012,\"Position\":26,\"Level\":2}]},\"Index\":9}"
//	//str := "{}"
//	strMap := make(map[string]interface{})
//	err := json.Unmarshal([]byte(str), &strMap)
//	if err != nil {
//		fmt.Printf("err=%+v\n", err)
//		return
//	}
//	fmt.Printf("strMap=%+v\n", strMap)
//
//	rpMap, ok := strMap["Replace"].(map[string]interface{})
//	if !ok {
//		fmt.Println("not ok")
//	}
//	fmt.Printf("rpMap[EarlyHeroes]=%+v\n", rpMap["EarlyHeroes"])
//	ehStr, err := json.Marshal(rpMap["EarlyHeroes"])
//	if err != nil {
//		fmt.Println("err=", err)
//		return
//	}
//	eMap := make([]map[string]interface{}, 0)
//	err = json.Unmarshal(ehStr, &eMap)
//	if err != nil {
//		fmt.Println("err=", err)
//		return
//	}
//
//	fmt.Printf("type = %T\n", eMap)
//	fmt.Printf("eMap = %+v\n", eMap)
//	rpMap["EarlyHeroes"] = eMap
//
//	//eMap, ok := rpMap["EarlyHeroes"].([]map[string]interface{})
//	//if !ok {
//	//	fmt.Println("not ok2")
//	//}
//
//	rules := make(map[string]interface{})
//	rules = map[string]interface{}{
//		"LineupGameType": "min=0,max=1",
//		"Index":          "min=0,max=9",
//		"Replace": map[string]interface{}{
//			"LordID":          "gte=3",
//			"Name":            "min=1",
//			"PreferredEquips": "gte=1",
//			"ID":              "gte=0,number",
//			"NewComer":        "eq=true|eq=false",
//			"Advance":         "eq=true|eq=false",
//			//"FinalHeroes":     "gt=0,dive,dive",
//
//			//"EarlyHeroes": "dive,gt=1",
//			"EarlyHeroes": map[string]interface{}{
//				"HeroID":   "min=0",
//				"Position": "min=11,max=48",
//				"Level":    "min=1,max=3",
//				"IsCenter": "|eq=true|eq=false",
//			},
//			//"FinalHeroes": map[string]interface{}{
//			//	"HeroID":   "min=0",
//			//	"EquipIDs": "min=0",
//			//	"Position": "min=11,max=48",
//			//	"Level":    "min=1,max=3",
//			//	"IsCenter": "eq=true|eq=false",
//			//},
//		},
//	}
//	val := validator.New()
//	got := val.ValidateMapCtx(context.Background(), strMap, rules)
//	fmt.Printf("got=%+v\n", got)
//
//}
