/**
 * @Author: boyyang
 * @Date: 2022-02-13 20:36:31
 * @LastEditTime: 2022-05-29 20:08:46
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\setupRouter\router.go
 */

package setupRouter

import (
	"blog/middleware"
	routers "blog/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	// 路由注册
	routers.LoginRouterInit(r) //登录注册
	r.StaticFS("/assets", http.Dir("./assets"))

	// 路由中间件 token验证
	r.Use(middleware.TokenVerification())

	routers.UserRouterInit(r)        //用户信息
	routers.AritcleRouterInit(r)     //文章博客
	routers.UploadRouterInit(r)      //文件上传
	routers.PictureWallRouterInit(r) //图片墙
	routers.TagsRouterInit(r)        //标签

	return r
}
