package email

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"net/smtp"
	"strings"
)

//发送邮件
//host smtp.exmail.qq.com
//sender support@trustex.club
//port 端口号 587
//password Youshiqingda110
//nickname 发送人昵称
//toEmails 收件人列表
//subject 主题
//body 内容
//contentType 类型
func Send(host string, port string, sender string, password string, nickname string, toEmails []string, subject string, body string, contentType string) error {
	//auth :=smtp.CRAMMD5Auth(sender, password)
	auth := smtp.PlainAuth("", sender, password, host)
	if contentType == "html" {
		contentType = "Content-Type: text/html; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain; charset=UTF-8"
	}
	msg := []byte("To: " + strings.Join(toEmails, ",") + "\r\nFrom: " + nickname +
		"<" + sender + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	err := smtp.SendMail(host+":"+port, auth, sender, toEmails, msg)

	if err != nil {
		masage := string(msg) + "\r\n" + err.Error()
		logs.Error(masage)
		return err
	}

	return nil
}

//一般发送方式
//nickname := "TET官方"
//contentType := "text"
func NormalSend(toEmails []string, nickname string, subject string, body string, contentType string) bool {
	host, _ := beego.AppConfig.String("email::host")
	port, _ := beego.AppConfig.String("email::port")
	sender, _ := beego.AppConfig.String("email::user")
	password, _ := beego.AppConfig.String("email::password")
	if contentType != "html" {
		contentType = "text"
	}
	err := Send(host, port, sender, password, nickname, toEmails, subject, body, contentType)
	if err == nil {
		return true
	}
	fmt.Println(err)
	return false
}
