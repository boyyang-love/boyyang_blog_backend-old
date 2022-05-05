/**
 * @Author: boyyang
 * @Date: 2022-04-03 00:35:57
 * @LastEditTime: 2022-04-30 10:23:18
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\controller\pictureWall\pictureWall.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package controller

import (
	"blog/models"
	"blog/setupDatabase"
	"blog/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 上传图片信息
func UploadPicture(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, _ := utils.ParseToken(token)
	var form models.PictureWall
	c.Bind(&form)
	form.UserID = claims.Id
	res := setupDatabase.DB.Create(&form)
	if res.Error == nil {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 1, Msg: "上传成功"}, map[string]interface{}{"id": form.ID}))
	} else {
		c.JSON(http.StatusBadRequest, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "上传失败"}, res.Error))
	}
}

// 更新图片信息
func UpdatePicture(c *gin.Context) {
	var form models.PictureWall
	var picture models.PictureWall
	c.Bind(&form)
	fmt.Println(form)
	form.Tags = []models.ImagesTag{{TagName: "手机壁纸"}}
	err := setupDatabase.
		DB.
		Debug().
		Omit("Author").
		Model(&picture).
		Update(&form).
		Error
	if err == nil {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 1, Msg: "更新成功"}, nil))
	} else {
		c.JSON(http.StatusBadRequest, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "更新失败"}, err))
	}
}

// 获取图片墙
func GetPicture(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	imagesTypes, _ := strconv.Atoi(c.Query("type"))
	imagesHidden, _ := strconv.Atoi(c.Query("hidden"))
	imagesStatus, _ := strconv.Atoi(c.Query("status"))

	var pictures []models.PictureWall
	var count int
	var err error
	if page == 0 && limit == 0 {
		err = setupDatabase.
			DB.
			Preload("Author").
			Where("type = ? and hidden = ? and status = ?", imagesTypes, imagesHidden, imagesStatus).
			Find(&pictures).
			Count(&count).
			Error
	} else {
		err = setupDatabase.
			DB.
			Limit(limit).
			Offset((page-1)*limit).
			Preload("Author").
			Where("type = ? and hidden = ? and status = ?", imagesTypes, imagesHidden, imagesStatus).
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

// 删除图片
func DeletePicture(c *gin.Context) {
	id := c.Param("id")
	var picture models.PictureWall
	err := setupDatabase.DB.First(&picture, id).Error
	if err == nil {
		err = setupDatabase.DB.Delete(&picture).Error
		if err == nil {
			c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 1, Msg: "删除成功"}, picture))
		} else {
			c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "删除失败"}, err))
		}
	} else {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "删除失败"}, err))
	}
}

// 获取图片详情
func GetPictureDetail(c *gin.Context) {
	id := c.Param("id")
	var picture models.PictureWall
	err := setupDatabase.DB.First(&picture, id).Error
	if err == nil {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 1, Msg: "获取成功"}, picture))
	} else {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "获取失败"}, err))
	}
}
