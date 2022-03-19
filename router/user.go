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
	"github.com/gin-gonic/gin"
)

func UserRouterInit(r *gin.Engine) {
	userRouters := r.Group("user")
	{
		userRouters.GET("/", func(c *gin.Context) {

			// userList := []models.User{}

			// setupDatabase.DB.Find(&userList)

			// c.JSON(200, gin.H{
			// 	"data": userList,
			// })

		})

		userRouters.GET("update", func(c *gin.Context) {
			c.String(200, "update user")
		})
	}
}
