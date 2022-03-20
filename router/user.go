/**
 * @Author: boyyang
 * @Date: 2022-02-13 20:33:06
 * @LastEditTime: 2022-02-15 11:25:49
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\websit\router\user.go
 */

package routers

import (
	"websit/controller"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(r *gin.Engine) {
	userRouters := r.Group("api")
	{
		userRouters.GET("updateUser", controller.UpdateUser)
	}
}
