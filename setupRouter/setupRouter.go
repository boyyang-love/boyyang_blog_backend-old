/**
 * @Author: boyyang
 * @Date: 2022-02-13 20:36:31
 * @LastEditTime: 2022-07-03 13:45:43
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\setupRouter\setupRouter.go
 */

package setupRouter

import (
	"blog/global"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// gin.DisableConsoleColor()
	// 日志文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	// 文件资源存储位置
	r.StaticFS("/assets", http.Dir("./assets"))
	// 中间件
	SetupRouterMiddleware(r)
	// 路由注册(不需要鉴权)
	SetupRouterWithoutAuth(r)
	// 记录日志
	SetupLoggerMiddeware(r)
	// 路由注册(需要鉴权)
	SetupRouterWithAuth(r)
	// 服务
	r.Run(global.Config.Servers.Port)
	return r
}

func LoggerWrite(r *gin.Engine) {
	panic("unimplemented")
}

func Logger() {
	panic("unimplemented")
}
