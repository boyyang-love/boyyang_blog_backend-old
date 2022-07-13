/**
 * @Author: boyyang
 * @Date: 2022-07-02 17:41:11
 * @LastEditTime: 2022-07-12 19:23:48
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\middleware\logger.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package middleware

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// 日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		claims, _ := utils.ParseToken(token)
		start_time := time.Now()
		c.Next()
		end_time := time.Now()
		log := models.Logger{
			Uid:           uint(claims.Id),
			Uri:           c.Request.URL.Path,
			StartTime:     start_time.Unix(),
			EndTime:       end_time.Unix(),
			RawQuery:      c.Request.URL.RawQuery,
			Host:          c.Request.Host,
			ClientIP:      c.ClientIP(),
			RequestMethod: c.Request.Method,
			ContentType:   c.Request.Header.Get("Content-Type"),
			Status:        c.Writer.Status(),
			UserAgent:     c.Request.UserAgent(),
		}
		global.DB.Create(&log)
	}
}
