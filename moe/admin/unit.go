package admin

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"path/filepath"
	"time"
)

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
	uuid := uuid.New()
	extension := filepath.Ext(originalFilename)
	uuidFilename := uuid.String() + extension
	now := time.Now()
	year := now.Year()
	month := now.Month()
	newFilename := filepath.Join("usr", "uploads", fmt.Sprintf("%d", year), fmt.Sprintf("%02d", month), uuidFilename)
	return newFilename
}
