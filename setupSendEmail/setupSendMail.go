/**
 * @Author: boyyang
 * @Date: 2022-06-10 11:01:13
 * @LastEditTime: 2022-07-08 10:19:31
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
	"net/smtp"

	"github.com/jordan-wright/email"
)

type EmailContent struct {
	Name string
	Urls []string
}

func SendEmail(to []string, title string, content EmailContent) (msg string, err error) {
	// 获取html文件
	t, err := template.ParseFiles("./setupSendEmail/email.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return "读取html文件失败", err
	}
	data := content
	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		return "模板渲染失败", err
	}
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
		return "邮件发送失败", err
	}
	return "邮件发送成功", nil
}
