/**
 * @Author: boyyang
 * @Date: 2022-04-03 09:04:46
 * @LastEditTime: 2022-06-10 15:41:13
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\router\pictureWall.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package routers

import (
	api "blog/api/pictureWall"

	"github.com/gin-gonic/gin"
)

func PictureWallRouterInit(r *gin.Engine) {
	pictureWallRouters := r.Group("api")
	{
		pictureWallRouters.POST("uploadPicture", api.UploadPicture)
		pictureWallRouters.GET("getPictures", api.GetPicture)
		pictureWallRouters.POST("updatePicture", api.UpdatePicture)
		pictureWallRouters.GET("deletePicture", api.DeletePicture)
	}
}
