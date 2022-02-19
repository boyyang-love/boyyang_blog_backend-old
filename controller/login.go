/*
 * @Author: boyyang
 * @Date: 2022-02-14 17:01:43
 * @LastEditTime: 2022-02-18 22:46:27
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\controller\login.go
 */

package controller

import (
	"websit/models"
	"websit/setupDatabase"
	"websit/utils"

	"github.com/gin-gonic/gin"
)

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
				utils.ReturnData(200, userMes, c)
			} else {
				utils.ReturnData(200, "密码错误", c)
			}
		} else {
			utils.ReturnData(200, "不存在该用户", c)
		}
	} else {
		utils.ReturnData(200, "用户名和密码为必填项", c)
	}

}

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
				utils.ReturnData(200, "用户注册成功", c)
			} else {
				utils.ReturnData(500, "数据库写入失败", c)
			}
		} else {
			utils.ReturnData(200, "该用户名已经注册", c)
		}
	} else {
		utils.ReturnData(200, "用户名和密码为必填项", c)
	}
}
