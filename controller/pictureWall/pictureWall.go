/**
 * @Author: boyyang
 * @Date: 2022-04-03 00:35:57
 * @LastEditTime: 2022-04-03 14:03:52
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\controller\pictureWall\pictureWall.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package controller

import (
	"net/http"
	"strconv"
	"websit/models"
	"websit/setupDatabase"
	"websit/utils"

	"github.com/gin-gonic/gin"
)

// 上传图片
func UploadPicture(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, _ := utils.ParseToken(token)
	var form models.PictureWall
	c.Bind(&form)
	id := claims.Id
	form.UserID = id
	err := setupDatabase.DB.Create(&form).Error
	if err == nil {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 1, Msg: "上传成功"}, form))
	}
}

// 获取图片墙
func GetPicture(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	var pictures []models.PictureWall
	var count int
	var err error
	if page == 0 && limit == 0 {
		err = setupDatabase.
			DB.
			Preload("Author").
			Find(&pictures).
			Count(&count).
			Error
	} else {
		err = setupDatabase.
			DB.
			Limit(limit).
			Offset((page - 1) * limit).
			Preload("Author").
			Find(&pictures).
			Limit(-1).
			Offset(-1).
			Count(&count).
			Error
	}
	if err == nil {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 1, Msg: "获取成功", Count: count}, pictures))
	} else {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "获取失败"}, err))
	}
}
