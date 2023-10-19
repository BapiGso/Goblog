package mymiddleware

import (
	"errors"
	"reflect"
)

// Struct2map 单层结构体转map,key是结构体成员的名称
func Struct2map(s any) (map[string]any, error) {
	v := reflect.ValueOf(s)
	// 如果是指针类型，获取指针指向的值
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// 检查是否是结构体类型
	if v.Kind() != reflect.Struct {
		return nil, errors.New("input is not a struct,maybe need input *struct")
	}
	t := reflect.TypeOf(s)

	data := make(map[string]any)
	for i := 0; i < v.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data, nil
}
