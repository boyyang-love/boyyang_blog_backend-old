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
	"fmt"
	"net/http"
	"path"
	"websit/models"
	"websit/setupDatabase"
	"websit/utils"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	token := c.Request.Header.Get("token")
	file, _ := c.FormFile("file")
	dst := path.Join("./assets", file.Filename)
	c.SaveUploadedFile(file, dst)
	claims, err := utils.ParseToken(token)
	fmt.Println(token)
	if err == nil {
		upload := models.Upload{
			Url:    dst,
			Name:   file.Filename,
			UserID: claims.Id,
		}
		err := setupDatabase.DB.Create(&upload).Error
		if err != nil {
			c.JSON(http.StatusOK, utils.RetunMsgFunc(utils.Code{Code: 0, Msg: "图片上传失败"}, err))
		} else {
			msg := map[string]interface{}{
				"name": file.Filename,
				"url":  "/assets/" + file.Filename,
			}
			fmt.Println(msg)
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
