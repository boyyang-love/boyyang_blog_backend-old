/**
 * @Author: boyyang
 * @Date: 2022-02-13 20:12:12
 * @LastEditTime: 2022-02-18 17:54:23
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\router\login.go
 */

package routers

import (
	"websit/controller"

	"github.com/gin-gonic/gin"
)

func LoginRouterInit(r *gin.Engine) {
	loginRouters := r.Group("api")
	{
		loginRouters.POST("login", controller.Login)
		loginRouters.POST("register", controller.Register)
	}
}
