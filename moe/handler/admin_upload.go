package handler

import (
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type UploadParam struct {
	Cid uint16 `query:"cid" form:"cid" json:"cid"`
}

// UploadImage 处理上传图片的请求 todo 多文件上传
func UploadImage(c echo.Context) error {
	// 从表单中获取上传的文件
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 读取文件的字节数据
	imgByte, err := io.ReadAll(src)
	if err != nil {
		return err
	}

	// 压缩图片
	compressedBytes := compressImageResource(imgByte)

	// 使用 UUID 和日期重命名文件
	renameFilePath := renameWithUUIDAndDate(file.Filename)

	// 创建文件路径中的所有目录
	err = os.MkdirAll(filepath.Dir(renameFilePath), 0755)
	if err != nil {
		return err
	}

	// 打开重命名后的文件
	f, err := os.OpenFile(renameFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	// 将压缩后的字节写入文件
	_, err = f.Write(compressedBytes)
	if err != nil {
		return err
	}

	// 返回成功响应，包括重命名后的文件路径、原始文件名作为 alt 和 title
	return c.JSON(http.StatusOK, struct {
		Url   string `json:"url"`
		Alt   string `json:"alt"`
		Title string `json:"title"`
	}{
		strings.Replace("/"+renameFilePath, "\\", "/", -1),
		file.Filename,
		file.Filename,
	})
}

func UploadTest(c echo.Context) error {
	return c.Render(200, "testupload.template", nil)
}
