/**
 * @Author: boyyang
 * @Date: 2022-02-16 17:25:46
 * @LastEditTime: 2022-04-05 14:45:12
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\router\upload.go
 */

package routers

import (
	controller "websit/controller/upload"

	"github.com/gin-gonic/gin"
)

func UploadRouterInit(r *gin.Engine) {
	uploadRouters := r.Group("api")
	{
		uploadRouters.POST("upload", controller.Upload)
		uploadRouters.GET("getImgs", controller.GetImgs)
		uploadRouters.GET("getAllImgs", controller.GetAllImgs)
	}
}
