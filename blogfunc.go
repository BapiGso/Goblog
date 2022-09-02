package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	highlighting "github.com/yuin/goldmark-highlighting"

	//"github.com/alecthomas/chroma/formatters/html"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"strconv"
	"strings"
	"time"
)

//类的结构可以按这个优化
//https://betterprogramming.pub/how-to-speed-up-your-struct-in-golang-76b846209587

var mon = map[string]string{
	"01": "一月",
	"02": "二月",
	"03": "三月",
	"04": "四月",
	"05": "五月",
	"06": "六月",
	"07": "七月",
	"08": "八月",
	"09": "九月",
	"10": "十月",
	"11": "十一月",
	"12": "十二月",
}

var sanitizer = bluemonday.StrictPolicy()
var mdoption = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		extension.Linkify,
		highlighting.NewHighlighting(
			highlighting.WithStyle("github")),
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
	goldmark.WithRendererOptions(
		html.WithHardWraps(),
		html.WithUnsafe(),
	),
)
var reCacheIndex, reCachePost chan bool

func init() {
	//go recache()
}

func recache() {
	//timeout := time.After(2 * time.Hour)
	for {
		select {
		case <-reCacheIndex:
			fmt.Println("index recazhe")
			return
		default:
			return
		}
	}
}

//时间戳转时间: 1660806518->十二月 23,2022
func unix2time(unix int64) string {
	tmp := (time.Unix(unix, 0)).Format("01 02, 2006")
	tmp = strings.Replace(tmp, tmp[:2], mon[tmp[:2]], 1)
	return tmp
}

//markdown转html
func md2html(source []byte) string {
	var buf bytes.Buffer
	if err := mdoption.Convert(source, &buf); err != nil {
		panic(err)
	}
	return buf.String()
}

//将html转text并且截取字数
func html2txt(html string) string {
	html = sanitizer.Sanitize(html)
	if len(html) <= 60 {
		return html
	} else {
		return html[:100]
	}
}

//首页返回1，不是数字返回err调用404，其他为对应页数
func isNum(numstr string) (uint64, error) {
	if numstr == "" {
		return 1, nil
	}
	num, err := strconv.ParseUint(numstr, 10, 64)
	if err != nil {
		return 0, err
	} else {
		return num, nil
	}
}

func hash(passwd string) string {
	h := sha1.New() // md5加密类似md5.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(passwd))
	bs := h.Sum(nil)
	h.Reset()
	passwdhash := hex.EncodeToString(bs) //转16进制
	return passwdhash
}

//todo random covermusic
//todo backup
