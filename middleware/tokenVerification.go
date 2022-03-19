/**
 * @Author: boyyang
 * @Date: 2022-02-18 15:29:14
 * @LastEditTime: 2022-02-19 11:17:03
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\middleware\tokenVerification.go
 */

package middleware

import (
	"net/http"
	"websit/utils"

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
