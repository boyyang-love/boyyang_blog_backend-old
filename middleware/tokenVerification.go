/**
 * @Author: boyyang
 * @Date: 2022-02-18 15:29:14
 * @LastEditTime: 2022-06-03 11:02:00
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\middleware\tokenVerification.go
 */

package middleware

import (
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenVerification() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header["Token"]

		if len(token) == 0 {
			c.JSON(
				http.StatusOK,
				utils.Msg(utils.Message{Code: 0, Msg: "token不能为空"}),
			)
			c.Abort()
			return
		} else {
			_, err := utils.ParseToken(token[0])
			if err != nil {
				c.JSON(
					http.StatusOK,
					utils.Msg(utils.Message{Code: 0, Msg: "token验证失败"}),
				)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
