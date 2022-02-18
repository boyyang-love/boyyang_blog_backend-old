/*
 * @Author: boyyang
 * @Date: 2022-02-18 15:29:14
 * @LastEditTime: 2022-02-18 16:35:36
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\middleware\tokenVerification.go
 */

package middleware

import (
	"websit/utils"

	"github.com/gin-gonic/gin"
)

func TokenVerification() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header["Token"]

		if len(token) == 0 {
			utils.ReturnData(100, "缺少token", c)
			c.Abort()
			return
		} else {
			_, err := utils.ParseToken(token[0])
			if err != nil {
				utils.ReturnData(205, "token 验证失败", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
