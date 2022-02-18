/*
 * @Author: boyyang
 * @Date: 2022-02-13 20:36:31
 * @LastEditTime: 2022-02-18 15:40:53
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\setupRouter\setupRouter.go
 */

package setupRouter

import (
	"net/http"
	"websit/middleware"
	routers "websit/router"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 路由注册
	routers.LoginRouterInit(r) //登录注册

	r.Use(middleware.TokenVerification())

	routers.UserRouterInit(r)    //用户信息
	routers.AritcleRouterInit(r) //文章博客
	routers.UploadRouterInit(r)  //文件上传

	r.StaticFS("/assets", http.Dir("./assets"))

	return r
}
