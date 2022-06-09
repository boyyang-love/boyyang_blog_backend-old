/**
 * @Author: boyyang
 * @Date: 2022-06-09 10:16:49
 * @LastEditTime: 2022-06-09 10:22:23
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\setupRouter\routerWithAuth.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package setupRouter

import (
	routers "blog/router"

	"github.com/gin-gonic/gin"
)

func SetupRouterWithAuth(r *gin.Engine) {
	routers.UserRouterInit(r)        //用户信息
	routers.AritcleRouterInit(r)     //文章博客
	routers.UploadRouterInit(r)      //文件上传
	routers.PictureWallRouterInit(r) //图片墙
	routers.TagsRouterInit(r)        //标签
}
