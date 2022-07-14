/**
 * @Author: boyyang
 * @Date: 2022-07-14 12:22:44
 * @LastEditTime: 2022-07-14 13:16:47
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\router\leaveMessage.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package routers

import (
	api "blog/api/leaveMessage"

	"github.com/gin-gonic/gin"
)

func LeaveMessageRouterInit(r *gin.Engine) {
	leaveMessageRouters := r.Group("api")
	{
		leaveMessageRouters.POST("/leaveMessage", api.LeaveMessageTo)
	}
}
