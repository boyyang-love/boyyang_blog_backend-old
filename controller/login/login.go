/**
 * @Author: boyyang
 * @Date: 2022-02-14 17:01:43
 * @LastEditTime: 2022-04-05 13:38:39
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\controller\login\login.go
 */

package controller

import (
	"net/http"
	"websit/models"
	"websit/setupDatabase"
	"websit/utils"

	"github.com/gin-gonic/gin"
)

// 登录
func Login(c *gin.Context) {
	username := c.PostForm("name")
	password := c.PostForm("password")
	if username != "" && password != "" {
		var user models.User
		err := setupDatabase.DB.Where("Username = ?", username).First(&user).Error
		if err == nil {
			if user.Password == utils.MD5(password) {
				token, _ := utils.GenerateToken(user.Username, user.Password, int(user.ID))
				userMes := map[string]interface{}{
					"info":  &user,
					"token": token,
				}
				c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 1, Msg: "登录成功"}, userMes))
			} else {
				c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "密码错误"}, nil))
			}
		} else {
			c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "不存在该用户"}, nil))
		}
	} else {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "账号和密码为必填项"}, nil))
	}

}

// 注册
func Register(c *gin.Context) {
	username := c.PostForm("name")
	password := c.PostForm("password")
	if username != "" && password != "" {
		var user models.User
		err := setupDatabase.DB.Where("Username = ?", username).First(&user).Error
		if err != nil {
			addUser := models.User{
				Username: username,
				Password: utils.MD5(password),
			}
			err := setupDatabase.DB.Create(&addUser).Error
			if err == nil {
				c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 1, Msg: "用户注册成功"}, nil))
			} else {
				c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "数据库写入失败"}, nil))
			}
		} else {
			c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "该用户名已经注册"}, nil))
		}
	} else {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "用户名和密码为必填项"}, nil))
	}
}
