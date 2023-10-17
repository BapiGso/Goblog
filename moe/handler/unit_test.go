package handler

import "testing"

func TestS_Struct2map(t *testing.T) {
	req := struct {
		Author string `xml:"author"   form:"author" validate:"required,min=1,max=200"`
		Mail   string `xml:"mail"     form:"mail" validate:"email,required,min=1,max=200"`
		Text   string `xml:"text"     form:"text" validate:"required,min=1,max=1000"`
		Url    string `xml:"url"      form:"url" validate:"omitempty,url,min=1,max=200" `
	}{
		"haha",
		"1@a.com",
		"hello",
		"a.com",
	}
	reqMap, _ := struct2map(req)
	t.Log(reqMap)
}
