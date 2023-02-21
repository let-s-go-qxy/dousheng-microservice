package test

import (
	"dousheng/kitex_gen/comment"
	"dousheng/kitex_gen/user"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
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

// Traversal 生成对应结构体忽略 json omitempty 的 map[string]interface，本质是复制结构体
// tips: 1.如果传入的是切片请给一个切片地址方便回传 2.注意给值的时候结构体内指针和切片不要给nil
func Traversal(theStruct interface{}, slice1 *[]interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	obj := reflect.Value{}
	// 解引用，同时考虑当前的theStruct是不是已经是reflect.Value了
	if reflect.ValueOf(theStruct).Kind().String() == "ptr" {
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
				*slice1 = append(*slice1, convertBaseValue(obj.Index(i), kind))
			} else if kind == "slice" {
				slice3 := &[]interface{}{}
				// 解引用
				if slice2.Index(i).Kind().String() == "ptr" {
					Traversal(slice2.Index(i).Elem(), slice3)
				} else {
					Traversal(slice2.Index(i), slice3)
				}
				*slice1 = append(*slice1, slice3)
			} else {
				// 解引用
				if slice2.Index(i).Kind().String() == "ptr" {
					*slice1 = append(*slice1, Traversal(slice2.Index(i).Elem(), nil))
				} else {
					*slice1 = append(*slice1, Traversal(slice2.Index(i), nil))
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
				res[convertName(name)] = Traversal(val, nil)
			} else if kind == "slice" {
				slice2 := &[]interface{}{}
				Traversal(val, slice2)
				res[convertName(name)] = slice2
			} else {
				res[convertName(name)] = convertBaseValue(val, kind)
			}
		}
	}
	return res
}

// 返回基础类型值
func convertBaseValue(val reflect.Value, kind string) interface{} {
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
func convertName(name string) string {
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

func TestTraversal(t *testing.T) {
	obj := comment.CommentListResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		CommentList: []*comment.Comment{
			{
				Id: 1,
				User: &user.User{
					Id:              1,
					Name:            "aei",
					FollowCount:     0,
					FollowerCount:   0,
					IsFollow:        false,
					WorkCount:       2,
					BackgroundImage: "img",
					Signature:       "",
					TotalFavorite:   1,
					FavoriteCount:   0,
					Avatar:          "",
				},
				Content:    "213",
				CreateDate: "133",
			},
			{
				Id:         0,
				User:       &user.User{},
				Content:    "",
				CreateDate: "",
			},
		},
	}
	t1 := Traversal(obj, nil)
	t2 := Traversal(&obj, nil)
	j1, _ := json.Marshal(t1)
	j2, _ := json.Marshal(t2)
	j3, _ := json.Marshal(obj)
	fmt.Println(string(j1), "\n------------\n",
		string(j2), "\n------------\n",
		string(j3))
	if string(j1) != string(j2) {
		t.Fail()
	}
}
