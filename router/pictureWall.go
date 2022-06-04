/**
 * @Author: boyyang
 * @Date: 2022-04-03 09:04:46
 * @LastEditTime: 2022-06-04 10:44:39
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\router\pictureWall.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package routers

import (
	controller "blog/controller/pictureWall"

	"github.com/gin-gonic/gin"
)

func PictureWallRouterInit(r *gin.Engine) {
	pictureWallRouters := r.Group("api")
	{
		pictureWallRouters.POST("uploadPicture", controller.UploadPicture)
		pictureWallRouters.GET("getPictures", controller.GetPicture)
		pictureWallRouters.POST("updatePicture", controller.UpdatePicture)
		pictureWallRouters.GET("deletePicture", controller.DeletePicture)
	}
}
