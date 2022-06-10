/**
 * @Author: boyyang
 * @Date: 2022-02-13 20:12:12
 * @LastEditTime: 2022-06-10 15:41:01
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\router\loginRegister.go
 */

package routers

import (
	api "blog/api/loginRegister"

	"github.com/gin-gonic/gin"
)

func LoginRouterInit(r *gin.Engine) {
	loginRegisterRouters := r.Group("api")
	{
		loginRegisterRouters.POST("login", api.Login)
		loginRegisterRouters.POST("register", api.Register)
	}
}
