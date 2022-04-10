/**
 * @Author: boyyang
 * @Date: 2022-02-14 17:01:53
 * @LastEditTime: 2022-04-10 21:49:51
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\controller\user\user.go
 */
package controller

import (
	"blog/models"
	"blog/setupDatabase"
	"blog/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	token := c.Request.Header.Get("token")
	name := c.PostForm("username")
	avaterUrl := c.PostForm("avaterUrl")
	age, _ := strconv.Atoi(c.PostForm("age"))
	sex, _ := strconv.Atoi(c.PostForm("sex"))
	birthday, _ := strconv.Atoi(c.PostForm("birthday"))
	userInfo, _ := utils.ParseToken(token)
	tel := c.PostForm("tel")
	email := c.PostForm("email")
	background := c.PostForm("background_img")
	var user models.User
	data := setupDatabase.
		DB.
		Where("Username = ?", name).
		First(&user)
	// 如果用户存
	if data.RowsAffected > 0 && user.ID != uint(userInfo.Id) {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "该用户名已经存在"}, nil))
		return
	}
	err := setupDatabase.
		DB.
		Model(models.User{}).
		Where("id = ?", userInfo.Id).
		Update(models.User{
			Username:   name,
			AvaterUrl:  avaterUrl,
			Age:        age,
			Sex:        sex,
			Birthday:   birthday,
			Background: background,
			TelPhone:   tel,
			Email:      email,
		}).
		Error
	if err == nil {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 1, Msg: "用户信息更新成功"}, nil))
	} else {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "用户信息更新失败"}, err))
	}
}
