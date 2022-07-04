/**
 * @Author: boyyang
 * @Date: 2022-07-02 18:02:51
 * @LastEditTime: 2022-07-02 18:07:11
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\setupRouter\loggerMiddeware.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */
package setupRouter

import (
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func SetupLoggerMiddeware(r *gin.Engine) {
	r.Use(middleware.Logger())
}
