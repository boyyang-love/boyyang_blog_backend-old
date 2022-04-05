/**
 * @Author: boyyang
 * @Date: 2022-02-18 15:29:14
 * @LastEditTime: 2022-04-05 15:39:35
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
			c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "缺少token"}, nil))
			c.Abort()
			return
		} else {
			_, err := utils.ParseToken(token[0])
			if err != nil {
				c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "token验证失败"}, nil))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
