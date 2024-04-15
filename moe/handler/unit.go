package handler

import (
	"SMOE/moe/database"
	"errors"
	"reflect"
)

func sortComms(data []database.Comments) [][]database.Comments {
	var final [][]database.Comments
	parentMap := make(map[uint]int)
	for _, v := range data {
		//父评论新建一个组，因为按时间排序肯定比子评论先
		if v.Parent == 0 {
			//初始化tmp的同时就把v加入切片
			tmp := []database.Comments{v}
			final = append(final, tmp)
			parentMap[v.Coid] = len(final) - 1
		} else { //如果是子评论，找自己属于哪个父评论组
			index := parentMap[v.Parent]
			final[index] = append(final[index], v)
			parentMap[v.Coid] = index
		}
	}
	//fmt.Println(parentIndexMap)
	return final
}

// struct2map 单层结构体转map,key是结构体成员的名称
func struct2map(s any) (map[string]any, error) {
	v := reflect.ValueOf(s)
	// 如果是指针类型，获取指针指向的值
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// 检查是否是结构体类型
	if v.Kind() != reflect.Struct {
		return nil, errors.New("input is not a struct,please input *struct")
	}
	t := reflect.TypeOf(s)

	data := make(map[string]any)
	for i := 0; i < v.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data, nil
}
