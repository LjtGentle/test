package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

type AnswerInvalidError struct {
	gojsonschema.ResultErrorFields
}

func newAnswerInvalidError(context *gojsonschema.JsonContext, value interface{}, details gojsonschema.ErrorDetails) *AnswerInvalidError {
	err := AnswerInvalidError{}
	err.SetContext(context)
	err.SetType("custom_invalid_error")
	// it is important to use SetDescriptionFormat() as this is used to call SetDescription() after it has been parsed
	// using the description of err will be overridden by this.
	err.SetDescriptionFormat("Answer to the Ultimate Question of Life, the Universe, and Everything is {{.answer}}")
	err.SetValue(value)
	err.SetDetails(details)

	return &err
}

func main() {
	test01()
}

func test01() {
	fmt.Println("test01")
	var formatSchema = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "type": "object",
    "properties": {
        "name": {
            "type": "string",
            "title": "名字(name)",
            "default": "张三",
            "minLength": 2,
            "maxLength": 4,
			"datasource": 1
        },
        "age": {
            "type": "number",
            "description": "这是说明这是说明",
            "minimum": 0,
            "maximum": 150
        },
        "locked": {
            "type": "boolean",
            "default": false
        },
        "tags": {
            "type": "array",
            "items": {
                "type": "string",
                "enum": [
                    "foor",
                    "bar"
                ]
            }
        },
        "game": {
            "type": "string",
            "format": "game:game123"
        },
        "qq_appid": {
            "type": "string",
            "format": "qq_appid"
        },
        "wx_appid": {
            "type": "string",
            "format": "wx_appid"
        }
    },
    "required": [
        "name",
        "age",
        "locked",
        "tags",
        "game",
        "qq_appid",
        "wx_appid"
    ],
    "title": "用户卡片"
}`
	gojsonschema.FormatCheckers.
		AddWithParam("game", gameChecker{}).
		Add("qq_appid", qqAppidChecker{}).
		Add("wx_appid", wxAppidChecker{})
	s1 := gojsonschema.NewStringLoader(formatSchema)
	validateResult, err := gojsonschema.Validate(s1, gojsonschema.NewGoLoader(map[string]interface{}{
		"name":     "张三",
		"age":      18,
		"locked":   true,
		"tags":     []string{"foor", "bar"},
		"game":     "smoba",
		"qq_appid": "123",
		"wx_appid": "abc",
	}))
	if err != nil {
		fmt.Printf("Validate err=%+v\n", err)
		return
	}

	if !validateResult.Valid() {
		for _, desc := range validateResult.Errors() {
			fmt.Printf("validateResult err=%+v\n", desc)
		}
		return
	}
	fmt.Println("success!")
}

type gameChecker struct {
}

var gameList = []string{"smoba", "qsm", "codm"}

//func (g gameChecker) IsFormat(input interface{}) bool {
//	game, ok := input.(string)
//	if !ok {
//		return false
//	}
//	for _, v := range gameList {
//		if game == v {
//			return true
//		}
//	}
//	return false
//}

func (g gameChecker) IsFormat(input interface{}, params []string) bool {
	game, ok := input.(string)
	if !ok {
		return false
	}
	fmt.Println("in gameChecker IsFormat params=", params)
	for _, v := range gameList {
		if game == v {
			return true
		}
	}
	return false
}

type qqAppidChecker struct {
}

var qqAppidList = []string{"123", "321", "1234567"}

func (g qqAppidChecker) IsFormat(input interface{}) bool {
	appid, ok := input.(string)
	if !ok {
		return false
	}
	for _, v := range qqAppidList {
		if appid == v {
			return true
		}
	}
	return false
}

type wxAppidChecker struct {
}

var wxAppidList = []string{"abc", "cba", "abcdefg"}

func (g wxAppidChecker) IsFormat(input interface{}) bool {
	appid, ok := input.(string)
	if !ok {
		return false
	}
	for _, v := range wxAppidList {
		if appid == v {
			return true
		}
	}
	return false
}
