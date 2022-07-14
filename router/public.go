/**
 * @Author: boyyang
 * @Date: 2022-07-14 13:15:22
 * @LastEditTime: 2022-07-14 13:17:18
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\router\public.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package routers

import (
	api "blog/api/public"

	"github.com/gin-gonic/gin"
)

func PublicRouterInit(r *gin.Engine) {
	publicRouters := r.Group("api")
	{
		publicRouters.GET("/public", api.Public)
	}
}
