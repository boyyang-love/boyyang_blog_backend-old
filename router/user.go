/**
 * @Author: boyyang
 * @Date: 2022-02-13 20:33:06
 * @LastEditTime: 2022-06-13 13:13:28
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\router\user.go
 */

package routers

import (
	api "blog/api/user"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(r *gin.Engine) {
	userRouters := r.Group("api")
	{
		userRouters.POST("updateUser", api.UpdateUser)
		userRouters.POST("updateUserPassword", api.UpdateUserPassword)
	}
}
