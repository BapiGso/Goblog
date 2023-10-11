package handler

import (
	"SMOE/moe/database"
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"path/filepath"
	"reflect"
	"strconv"
	"time"
)

// validateNum 首页返回1，不是数字返回err调用404，其他为对应页数
func validateNum(numStr string) (int, error) {
	if numStr == "" {
		return 1, nil
	}
	num, err := strconv.ParseUint(numStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(num), nil
}

func sortComms(data []database.Comments) [][]database.Comments {
	var final [][]database.Comments
	parentMap := make(map[uint32]int)
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

// struct2map 单层结构体转map
func struct2map(s any) map[string]any {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	data := make(map[string]any)
	for i := 0; i < v.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func compressImageResource(data []byte) []byte {
	imgSrc, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return data
	}
	newImg := image.NewRGBA(imgSrc.Bounds())
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)

	// 使用buffer来存储压缩后的图像数据
	buf := new(bytes.Buffer)

	// 压缩图像，并将压缩后的数据写入buffer
	err = jpeg.Encode(buf, newImg, &jpeg.Options{Quality: 75})
	if err != nil {
		return data
	}

	// 如果压缩后的数据比原始数据更大，则返回原始数据
	if buf.Len() >= len(data) {
		return data
	}

	// 返回压缩后的数据
	return buf.Bytes()
}

func renameWithUUIDAndDate(originalFilename string) string {
	uuider := uuid.New()
	extension := filepath.Ext(originalFilename)
	uuidFilename := uuider.String() + extension
	now := time.Now()
	year := now.Year()
	month := now.Month()
	newFilename := filepath.Join("usr", "uploads", fmt.Sprintf("%d", year), fmt.Sprintf("%02d", month), uuidFilename)
	return newFilename
}
