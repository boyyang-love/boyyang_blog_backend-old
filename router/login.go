/**
 * @Author: boyyang
 * @Date: 2022-02-13 20:12:12
 * @LastEditTime: 2022-06-03 11:21:36
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\router\login.go
 */

package routers

import (
	controller "blog/controller/loginRegister"

	"github.com/gin-gonic/gin"
)

func LoginRouterInit(r *gin.Engine) {
	loginRouters := r.Group("api")
	{
		loginRouters.POST("login", controller.Login)
		loginRouters.POST("register", controller.Register)
	}
}
