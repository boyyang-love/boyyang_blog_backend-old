/**
 * @Author: boyyang
 * @Date: 2022-04-29 11:20:12
 * @LastEditTime: 2022-04-29 11:23:54
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\router\tags.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package routers

import (
	controller "blog/controller/tags"

	"github.com/gin-gonic/gin"
)

func TagsRouterInit(r *gin.Engine) {
	tagsRouters := r.Group("api")
	{
		tagsRouters.POST("tags/add", controller.AddTags)
	}
}
