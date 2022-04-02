/**
 * @Author: boyyang
 * @Date: 2022-02-16 17:27:10
 * @LastEditTime: 2022-02-19 11:59:10
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\controller\upload.go
 */

package controller

import (
	"context"
	"net/http"
	"websit/client"
	"websit/models"
	"websit/setupDatabase"
	"websit/utils"

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
			Url:    "/images/" + file.Filename,
			Name:   file.Filename,
			UserID: claims.Id,
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
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 1, Msg: "获取失败"}, nil))
	} else {
		c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 1, Msg: "获取成功", Count: count}, imgs))
	}
}
