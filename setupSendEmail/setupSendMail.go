/**
 * @Author: boyyang
 * @Date: 2022-06-10 11:01:13
 * @LastEditTime: 2022-06-10 19:43:24
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\setupSendEmail\setupSendMail.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package setupSendEmail

import (
	"blog/global"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmail(to string, title string, content string) {
	emailConfig := global.Config.Email
	e := &email.Email{
		From:    emailConfig.From,
		To:      []string{to},
		Subject: title,
		Text:    []byte(content),
	}
	//设置服务器相关的配置
	err := e.Send(
		"smtp.qq.com:25",
		smtp.PlainAuth("", emailConfig.From, emailConfig.SmtpKey, "smtp.qq.com"),
	)
	if err != nil {
		log.Fatal(err)
	}
}
