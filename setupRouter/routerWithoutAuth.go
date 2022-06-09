/**
 * @Author: boyyang
 * @Date: 2022-06-09 10:17:04
 * @LastEditTime: 2022-06-09 13:08:16
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\setupRouter\routerWithoutAuth.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package setupRouter

import (
	routers "blog/router"

	"github.com/gin-gonic/gin"
)

func SetupRouterWithoutAuth(r *gin.Engine) {
	routers.LoginRouterInit(r) //登录注册
}
