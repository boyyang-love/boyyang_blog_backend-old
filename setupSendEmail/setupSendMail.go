/**
 * @Author: boyyang
 * @Date: 2022-06-10 11:01:13
 * @LastEditTime: 2022-06-10 15:46:17
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\setupSendEmail\setupSendMail.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package setupSendEmail

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmail(to string, title string, content string) {
	// template
	e := email.NewEmail()
	e.From = "boyyang个人博客网站 <1761617270@qq.com>"
	// 设置接收方的邮箱
	e.To = []string{to}
	//设置主题
	e.Subject = title
	//设置文件发送的内容
	e.Text = []byte(content)
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "1761617270@qq.com", "nqhwetnrlyxxdedi", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
}
