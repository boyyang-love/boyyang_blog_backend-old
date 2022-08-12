/**
 * @Author: boyyang
 * @Date: 2022-06-09 10:16:49
 * @LastEditTime: 2022-08-10 16:54:28
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog_server\setupRouter\routerWithAuth.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package setupRouter

import (
	"blog/middleware"
	routers "blog/router"

	"github.com/gin-gonic/gin"
)

func SetupRouterWithAuth(r *gin.Engine) {
	// 路由中间件 token验证
	r.Use(middleware.TokenVerification())
	routers.UserRouterInit(r)         //用户信息
	routers.AritcleRouterInit(r)      //文章博客
	routers.UploadRouterInit(r)       //文件上传
	routers.PictureWallRouterInit(r)  //图片墙
	routers.TagsRouterInit(r)         //标签
	routers.SocketRouterInit(r)       //websocket
	routers.LeaveMessageRouterInit(r) //留言
	routers.PublicRouterInit(r)
}
