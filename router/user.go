/**
 * @Author: boyyang
 * @Date: 2022-02-13 20:33:06
 * @LastEditTime: 2022-06-10 15:42:01
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\router\user.go
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
	}
}
