/**
 * @Author: boyyang
 * @Date: 2022-06-10 14:57:50
 * @LastEditTime: 2022-06-10 15:50:15
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\router\sendEmail.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package routers

import (
	api "blog/api/sendEmail"

	"github.com/gin-gonic/gin"
)

func SendMailRouterInit(r *gin.Engine) {
	sendMailRouters := r.Group("api")
	{
		sendMailRouters.GET("sendEmail", api.SendEmail)
	}
}
