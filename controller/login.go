/*
 * @Author: boyyang
 * @Date: 2022-02-14 17:01:43
 * @LastEditTime: 2022-02-18 17:27:38
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

	// 写入一条用户
	// user := models.User{
	// 	Username:  "boyyang-love",
	// 	Password:  "1234567",
	// 	AvaterUrl: "url:jjjjjj",
	// 	Age:       18,
	// 	Sex:       "man",
	// }
	// setupDatabase.DB.Create(&user)

	// 写入一条文章
	// articel := models.Article{
	// 	Title:    "JavaScript学习教程",
	// 	Subtitle: "JavaScript",
	// 	Content:  "涛涛涛涛",
	// 	UserID:   1,
	// }
	// setupDatabase.DB.Create(&articel)

	// 关联查询 一条
	// var list models.Article
	// setupDatabase.DB.Preload("Author").First(&list)

	// 查询所有 user
	var list []models.Article
	setupDatabase.DB.Preload("Author").Find(&list)

	c.JSON(200, gin.H{
		"code":    200,
		"message": list,
	})
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
