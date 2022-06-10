/**
 * @Author: boyyang
 * @Date: 2022-02-16 17:25:46
 * @LastEditTime: 2022-06-10 15:41:50
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\router\upload.go
 */

package routers

import (
	api "blog/api/upload"

	"github.com/gin-gonic/gin"
)

func UploadRouterInit(r *gin.Engine) {
	uploadRouters := r.Group("api")
	{
		uploadRouters.POST("upload", api.Upload)
		uploadRouters.GET("getImgs", api.GetImgs)
		uploadRouters.GET("getAllImgs", api.GetAllImgs)
		uploadRouters.GET("deleteUpload", api.DeleteUpload)
	}
}
