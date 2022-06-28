/**
 * @Author: boyyang
 * @Date: 2022-06-10 11:01:13
 * @LastEditTime: 2022-06-28 17:14:14
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\setupSendEmail\setupSendMail.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package setupSendEmail

import (
	"blog/global"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type EmailContent struct {
	Name string
	Urls []string
}

func SendEmail(to []string, title string, content EmailContent) {
	// 获取html文件
	t, err := template.ParseFiles("./setupSendEmail/email.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	data := content
	buf := new(bytes.Buffer)
	t.Execute(buf, data)
	emailConfig := global.Config.Email
	e := &email.Email{
		From:    emailConfig.From,
		To:      to,
		Subject: title,
		HTML:    buf.Bytes(),
	}
	//设置服务器相关的配置
	err = e.Send(
		"smtp.qq.com:25",
		smtp.PlainAuth("", emailConfig.From, emailConfig.SmtpKey, "smtp.qq.com"),
	)
	if err != nil {
		log.Fatal(err)
	}
}
