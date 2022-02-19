/*
 * @Author: boyyang
 * @Date: 2022-02-16 17:25:46
 * @LastEditTime: 2022-02-19 11:47:09
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\router\upload.go
 */

package routers

import (
	"websit/controller"

	"github.com/gin-gonic/gin"
)

func UploadRouterInit(r *gin.Engine) {
	uploadRouters := r.Group("upload")
	{
		uploadRouters.POST("/", controller.Upload)
		uploadRouters.GET("/getImgs", controller.GetImgs)
	}
}
