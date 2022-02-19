/*
 * @Author: boyyang
 * @Date: 2022-02-13 20:36:31
 * @LastEditTime: 2022-02-19 11:39:07
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
	r.StaticFS("/assets", http.Dir("./assets"))

	// 路由中间件
	r.Use(middleware.TokenVerification())

	routers.UserRouterInit(r)    //用户信息
	routers.AritcleRouterInit(r) //文章博客
	routers.UploadRouterInit(r)  //文件上传

	return r
}
