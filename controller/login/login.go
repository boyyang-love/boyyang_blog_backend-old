/**
 * @Author: boyyang
 * @Date: 2022-02-14 17:01:43
 * @LastEditTime: 2022-06-03 10:45:39
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\controller\login\login.go
 */

package controller

import (
	"blog/models"
	"blog/setupDatabase"
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
		var user models.User
		res := setupDatabase.
			DB.
			Where("Username = ?", username).
			First(&user)
		if res.RowsAffected != 0 {
			if user.Password == utils.MD5(password) {
				token, _ := utils.GenerateToken(user.Username, user.Password, int(user.ID))
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

// 注册
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if strings.Trim(username, "") != "" && strings.Trim(password, "") != "" {
		var user models.User
		res := setupDatabase.
			DB.
			Where("Username = ?", username).
			First(&user)
		if res.RowsAffected == 0 {
			addUser := models.User{
				Username: username,
				Password: utils.MD5(password),
			}
			err := setupDatabase.
				DB.
				Create(&addUser).
				Error
			if err == nil {
				c.JSON(
					http.StatusOK,
					utils.Msg(utils.Message{Code: 1, Msg: "注册成功"}),
				)
			} else {
				c.JSON(
					http.StatusBadRequest,
					utils.Msg(utils.Message{Code: 0, Msg: "注册失败"}),
				)
			}
		} else {
			c.JSON(
				http.StatusBadRequest,
				utils.Msg(utils.Message{Code: 0, Msg: "该用户已存在"}),
			)
		}
	} else {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "账号和密码为必填项"}),
		)
	}
}
