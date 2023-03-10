package mail

import (
	"crypto/tls"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
	"net/smtp"
)

type Email struct {
	user   string
	passwd string
	Auth   smtp.Auth
	Conn   *tls.Conn
}

func (e *Email) Login(Db sqlx.DB) {
	var data []byte
	_ = Db.QueryRow(`SELECT value
		FROM typecho_options
		WHERE name='plugin:GoMail' `).Scan(&data)
	err := json.Unmarshal(data, &Email{})
	if err != nil {
		log.Println("获取邮箱用户名和密码失败")
	}
	e.Auth = smtp.PlainAuth("", e.user, e.passwd, "smtp.gmail.com")
	smtpServer := "smtp.gmail.com"
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer,
	}
	e.Conn, err = tls.Dial("tcp", smtpServer+":587", tlsConfig)
	if err != nil {
		log.Fatal(err)
	}
}
