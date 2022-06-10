/**
 * @Author: boyyang
 * @Date: 2022-02-16 17:27:10
 * @LastEditTime: 2022-06-10 15:39:59
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\api\upload\upload.go
 */

package api

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func Upload(c *gin.Context) {
	token := c.Request.Header.Get("token")
	file, _ := c.FormFile("file")
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "image/jpeg",
		},
	}
	claims, err := utils.ParseToken(token)
	if err == nil {
		path := fmt.Sprintf("/%d/%s/%s", claims.Id, "images", file.Filename)
		f, _ := file.Open()
		var err error
		_, err = global.Client.Object.Put(context.Background(), path, f, opt)
		if err != nil {
			panic(err)
		}
		upload := models.Upload{
			Url:      path,
			FileName: file.Filename,
			UserID:   claims.Id,
		}
		err = global.
			DB.
			Create(&upload).
			Error
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				utils.Msg(utils.Message{Code: 0, Msg: "上传失败", Error: err}),
			)
		} else {
			msg := map[string]interface{}{
				"fileName": file.Filename,
				"url":      path,
			}
			c.JSON(
				http.StatusOK,
				utils.Msg(utils.Message{Code: 1, Msg: "上传成功", Data: msg}),
			)
		}
	} else {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "上传失败", Error: err}),
		)
	}
}

func GetImgs(c *gin.Context) {
	id := c.Query("id")
	userId := c.Query("user_id")
	imgs := []models.Upload{}
	var err error
	var count int
	if id != "" {
		err = global.
			DB.
			Preload("Author").
			Where("id = ?", id).First(&imgs).
			Count(&count).
			Error
	} else if userId != "" {
		err = global.
			DB.
			Preload("Author").
			Where("upload_id = ?", userId).
			Find(&imgs).
			Count(&count).
			Error
	} else {
		err = global.
			DB.
			Preload("Author").
			Find(&imgs).
			Count(&count).
			Error
	}
	if err != nil {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 0, Msg: "获取失败", Error: err}),
		)
	} else {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 1, Msg: "获取成功", Count: count, Data: imgs}),
		)
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
		err = global.
			DB.
			Order("id desc").
			Preload("Author").
			Find(&imgs).
			Count(&count).
			Error
	} else {
		err = global.
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
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 0, Msg: "获取失败", Error: err}),
		)
	} else {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 1, Msg: "获取成功", Count: count, Data: imgs}),
		)
	}
}
