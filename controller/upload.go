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

	if err == nil {
		upload := models.Upload{
			Url:      dst,
			Name:     file.Filename,
			UploadID: claims.Id,
		}

		err := setupDatabase.DB.Create(&upload).Error

		if err != nil {
			// utils.ReturnData(200, err, c)
		} else {
			msg := map[string]interface{}{
				"name": file.Filename,
				"url":  "/assets/" + file.Filename,
			}

			fmt.Println(msg)

			// utils.ReturnData(200, msg, c)
		}
	} else {
		// utils.ReturnData(200, "token解析失败", c)
	}
}

func GetImgs(c *gin.Context) {
	id := c.Query("id")
	userId := c.Query("user_id")
	imgs := []models.Upload{}

	var err error

	if id != "" {
		err = setupDatabase.DB.Where("id = ?", id).First(&imgs).Error
	} else if userId != "" {
		err = setupDatabase.DB.Where("upload_id = ?", userId).Find(&imgs).Error
	} else {
		err = setupDatabase.DB.Find(&imgs).Error
	}

	if err != nil {
		// utils.ReturnData(200, err, c)
	} else {
		// utils.ReturnData(200, imgs, c)
	}

}
