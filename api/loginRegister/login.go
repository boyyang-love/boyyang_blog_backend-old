/**
 * @Author: boyyang
 * @Date: 2022-02-14 17:01:43
 * @LastEditTime: 2022-06-23 18:41:49
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\loginRegister\login.go
 */

package api

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 登录
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if strings.Trim(username, "") != "" && strings.Trim(password, "") != "" {
		res := global.
			DB.
			Where("Username = ?", username).
			Find(&models.User{})
		if res.RowsAffected != 0 {
			var user models.User
			res.Scan(&user)
			if user.Password == utils.MD5(password) {
				token, _ := utils.GenerateToken(user.Username, user.Password, int(user.ID), *user.Email)
				userMes := map[string]interface{}{
					"info":  &user,
					"token": token,
				}
				c.JSON(
					http.StatusOK,
					utils.Msg(utils.Message{Code: 1, Msg: "登录成功", Data: userMes}),
				)
			} else {
				c.JSON(
					http.StatusBadRequest,
					utils.Msg(utils.Message{Code: 0, Msg: "密码错误"}),
				)
			}
		} else {
			c.JSON(
				http.StatusBadRequest,
				utils.Msg(utils.Message{Code: 0, Msg: "不存在该用户"}),
			)
		}
	} else {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "账号和密码为必填项"}),
		)
	}
}
