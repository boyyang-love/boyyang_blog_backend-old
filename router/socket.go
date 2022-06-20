/**
 * @Author: boyyang
 * @Date: 2022-06-14 10:46:32
 * @LastEditTime: 2022-06-14 11:11:59
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\router\socket.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package routers

import (
	api "blog/api/socket"

	"github.com/gin-gonic/gin"
)

func SocketRouterInit(r *gin.Engine) {
	socketRouters := r.Group("api")
	{
		socketRouters.GET("socket", api.Socket)
	}
}
