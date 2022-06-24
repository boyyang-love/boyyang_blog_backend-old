/**
 * @Author: boyyang
 * @Date: 2022-02-13 20:36:31
 * @LastEditTime: 2022-06-24 18:13:35
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\setupRouter\setupRouter.go
 */

package setupRouter

import (
	"blog/global"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// gin.DisableConsoleColor()
	// 日志文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	// 文件资源存储位置
	r.StaticFS("/assets", http.Dir("./assets"))
	// 中间件
	SetupRouterMiddleware(r)
	// 路由注册(不需要鉴权)
	SetupRouterWithoutAuth(r)
	// 路由注册(需要鉴权)
	SetupRouterWithAuth(r)
	fmt.Println("💞🎈🎐路由初始化成功")
	// 服务
	r.Run(global.Config.Servers.Port)
	return r
}
