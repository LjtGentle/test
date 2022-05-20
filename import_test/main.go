package main

import (
	"fmt"
	"github.com/epiclabs-io/elastic"
	"reflect"
)

func main() {
	type stu struct {
		age  int
		name string
	}
	s1 := stu{
		age:  10,
		name: "tom",
	}
	s2 := stu{
		age:  10,
		name: "tom",
	}
	s := make([]stu, 0, 10)
	s = append(s, s1)
	s = append(s, s2)
	ty := ElementType()
	convert, err := elastic.Convert(s, ty)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Printf("convert =%+v\n", convert)
}

func ElementType(v interface{}) reflect.Type {
	rv := reflect.TypeOf(v)
	if rv.Kind() == reflect.Slice {
		return rv.Elem()
	}
	return rv
}


