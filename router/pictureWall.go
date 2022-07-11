/**
 * @Author: boyyang
 * @Date: 2022-04-03 09:04:46
 * @LastEditTime: 2022-07-11 13:07:31
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\router\pictureWall.go
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
		pictureWallRouters.GET("getPictures", api.GetPictures)
		pictureWallRouters.POST("updatePicture", api.UpdatePicture)
		pictureWallRouters.GET("deletePicture", api.DeletePicture)
		pictureWallRouters.GET("thumbsUpPicture", api.ThumbsUpPicture)
	}
}
