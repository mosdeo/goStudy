package main

import (
	"fmt"
	"reflect"
)

type TagType struct { // tags
	field1 bool   "==This is field1 tag=="
	field2 string `==This is field2 tag==`
	field3 int    "==This is field3 tag=="
}

func main() {
	//tt := TagType{true, "這是字串的內容", 1}
	tt := TagType{} //Don't care content
	for i := 0; i < 3; i++ {
		structField := GetFieldByMemberIndex(tt, i)
		fmt.Printf("field%d's tag: %v, name: %v\n", i+1, structField.Tag, structField.Name)
	}
}

func GetFieldByMemberIndex(tt TagType, index int) reflect.StructField {
	ttType := reflect.TypeOf(tt) //取得該 struct 的型別
	ixField := ttType.Field(index)
	return ixField
}
