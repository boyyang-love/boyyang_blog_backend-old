/**
 * @Author: boyyang
 * @Date: 2022-02-13 20:33:06
 * @LastEditTime: 2022-04-05 15:39:10
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\router\user.go
 */

package routers

import (
	controller "blog/controller/user"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(r *gin.Engine) {
	userRouters := r.Group("api")
	{
		userRouters.POST("updateUser", controller.UpdateUser)
	}
}
