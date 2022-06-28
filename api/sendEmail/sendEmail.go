/**
 * @Author: boyyang
 * @Date: 2022-06-10 14:50:28
 * @LastEditTime: 2022-06-28 17:04:18
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\sendEmail\sendEmail.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package api

import (
	"blog/setupSendEmail"

	"github.com/gin-gonic/gin"
)

func SendEmail(c *gin.Context) {
	to := []string{"1761617270@qq.com"}
	title := "个人博客网站注册"
	setupSendEmail.SendEmail(to, title, setupSendEmail.EmailContent{Name: "小明"})
}
