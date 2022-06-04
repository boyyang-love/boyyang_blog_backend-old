/**
 * @Author: boyyang
 * @Date: 2022-02-13 20:12:12
 * @LastEditTime: 2022-06-03 11:25:37
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\router\loginRegister.go
 */

package routers

import (
	controller "blog/controller/loginRegister"

	"github.com/gin-gonic/gin"
)

func LoginRouterInit(r *gin.Engine) {
	loginRegisterRouters := r.Group("api")
	{
		loginRegisterRouters.POST("login", controller.Login)
		loginRegisterRouters.POST("register", controller.Register)
	}
}
