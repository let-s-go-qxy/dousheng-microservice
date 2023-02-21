package utils

import (
	"fmt"
	"reflect"
	"strings"
)

var DefaultField = map[string]bool{
	"state":         true,
	"sizeCache":     true,
	"unknownFields": true,
	"flag":          true,
	"ptr":           true,
	"typ":           true,
}
var BaseType = map[string]bool{
	"int":    true,
	"int32":  true,
	"int64":  true,
	"string": true,
	"bool":   true,
}

// ConvertStruct 生成对应结构体忽略 json omitempty 的 map[string]interface，本质是复制结构体
// tips: 1.如果传入的是切片请给一个切片地址方便回传 2.注意给值的时候结构体内指针和切片不要给nil
func ConvertStruct(theStruct interface{}, slice1 *[]interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	obj := reflect.Value{}
	// 解引用，同时考虑当前的theStruct是不是已经是reflect.Value了
	if reflect.ValueOf(theStruct).Kind().String() == "ptr" {
		// 防止空值，但是这里应该不从我这处理，此处为兜底
		if reflect.ValueOf(theStruct).IsNil() {
			fmt.Println("[error]------------->", "传值为空")
			return res
		}
		if reflect.ValueOf(theStruct).Type().String() == "reflect.Value" {
			obj = theStruct.(reflect.Value).Elem()
		} else {
			obj = reflect.ValueOf(theStruct).Elem()
		}
	} else {
		if reflect.ValueOf(theStruct).Type().String() == "reflect.Value" {
			obj = theStruct.(reflect.Value)
		} else {
			obj = reflect.ValueOf(theStruct)
		}
	}
	// 如果传进来的是切片，对其进行预处理 TODO 这里并不是很严谨，问题是怎么判断
	if slice1 != nil {
		slice2 := obj
		for i := 0; i < slice2.Len(); i++ {
			kind := reflect.TypeOf(slice2.Index(i)).Kind().String()
			if BaseType[kind] {
				*slice1 = append(*slice1, ConvertBaseValue(obj.Index(i), kind))
			} else if kind == "slice" {
				slice3 := &[]interface{}{}
				// 解引用
				if slice2.Index(i).Kind().String() == "ptr" {
					ConvertStruct(slice2.Index(i).Elem(), slice3)
				} else {
					ConvertStruct(slice2.Index(i), slice3)
				}
				*slice1 = append(*slice1, slice3)
			} else {
				// 解引用
				if slice2.Index(i).Kind().String() == "ptr" {
					*slice1 = append(*slice1, ConvertStruct(slice2.Index(i).Elem(), nil))
				} else {
					*slice1 = append(*slice1, ConvertStruct(slice2.Index(i), nil))
				}
			}
		}
		return res
	}
	// reflect.Value 进行处理
	objType := obj.Type()
	for i := 0; i < obj.NumField(); i++ {
		name := objType.Field(i).Name
		// 忽略 kitex 自带字段
		if !DefaultField[name] {
			kind := obj.FieldByName(name).Kind().String()
			val := obj.FieldByName(name)
			// 解引用
			if kind == "ptr" {
				kind = obj.FieldByName(name).Elem().Kind().String()
				val = val.Elem()
			}
			if kind == "struct" {
				res[ConvertName(name)] = ConvertStruct(val, nil)
			} else if kind == "slice" {
				slice2 := &[]interface{}{}
				ConvertStruct(val, slice2)
				res[ConvertName(name)] = slice2
			} else {
				res[ConvertName(name)] = ConvertBaseValue(val, kind)
			}
		}
	}
	return res
}

// 返回基础类型值
func ConvertBaseValue(val reflect.Value, kind string) interface{} {
	switch kind {
	case "int", "int32", "int64":
		return val.Int()
	case "string":
		return val.String()
	case "bool":
		return val.Bool()
	}
	return nil
}

// 返回名字 大驼峰 -> 蛇形
func ConvertName(name string) string {
	data := make([]byte, 0, len(name)*2)
	j := false
	num := len(name)
	for i := 0; i < num; i++ {
		d := name[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}
