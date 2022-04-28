/**
 * @Author: boyyang
 * @Date: 2022-02-16 17:27:10
 * @LastEditTime: 2022-04-23 16:52:26
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\controller\upload\upload.go
 */

package controller

import (
	"blog/client"
	"blog/models"
	"blog/setupDatabase"
	"blog/utils"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func Upload(c *gin.Context) {
	token := c.Request.Header.Get("token")
	file, _ := c.FormFile("file")
	client := client.SetupClient()
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "image/jpeg",
		},
	}

	f, _ := file.Open()
	_, err := client.Object.Put(context.Background(), "/images/"+file.Filename, f, opt)
	if err != nil {
		panic(err)
	}
	claims, err := utils.ParseToken(token)
	if err == nil {
		upload := models.Upload{
			Url:      "/images/" + file.Filename,
			FileName: file.Filename,
			UserID:   claims.Id,
		}
		err := setupDatabase.DB.Create(&upload).Error
		if err != nil {
			c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "图片上传失败"}, err))
		} else {
			msg := map[string]interface{}{
				"name": file.Filename,
				"url":  "/images/" + file.Filename,
			}
			c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 1, Msg: "图片上传成功"}, msg))
		}
	} else {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "token验证失败"}, nil))
	}
}

func GetImgs(c *gin.Context) {
	id := c.Query("id")
	userId := c.Query("user_id")
	imgs := []models.Upload{}
	var err error
	var count int
	if id != "" {
		err = setupDatabase.DB.Preload("Author").Where("id = ?", id).First(&imgs).Count(&count).Error
	} else if userId != "" {
		err = setupDatabase.DB.Preload("Author").Where("upload_id = ?", userId).Find(&imgs).Count(&count).Error
	} else {
		err = setupDatabase.DB.Preload("Author").Find(&imgs).Count(&count).Error
	}
	if err != nil {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "获取失败"}, nil))
	} else {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 1, Msg: "获取成功", Count: count}, imgs))
	}
}

// 获取所有图片
func GetAllImgs(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	imgs := []models.Upload{}
	var count int
	var err error
	if page == 0 && limit == 0 {
		err = setupDatabase.
			DB.
			Order("id desc").
			Preload("Author").
			Find(&imgs).
			Count(&count).
			Error
	} else {
		err = setupDatabase.
			DB.
			Order("id desc").
			Limit(limit).
			Offset((page - 1) * limit).
			Preload("Author").
			Find(&imgs).
			Limit(-1).
			Offset(-1).
			Count(&count).
			Error
	}
	if err != nil {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "获取失败"}, nil))
	} else {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 1, Msg: "获取成功", Count: count}, imgs))
	}
}
