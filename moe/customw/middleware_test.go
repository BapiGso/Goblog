package customw

import (
	"testing"
)

func Test_setDefaults(t *testing.T) {
	type person struct {
		Name   string `default:"haha"`
		Age    int    `default:"17"`
		Weight int    `default:"50"`
	}
	p := new(person)
	setDefaults(p)
	if p.Name != "haha" && p.Age != 17 {
		t.Error("绑定默认值失败")
	}
}

//func BenchmarkBrotliWithConfig(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		brotli.NewWriterOptions(nil, brotli.WriterOptions{Quality: 11})
//	}
//}
