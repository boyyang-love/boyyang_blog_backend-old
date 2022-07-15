/**
 * @Author: boyyang
 * @Date: 2022-06-03 11:18:03
 * @LastEditTime: 2022-07-15 17:02:56
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\pictureWall\getPictures.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */
package api

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取图片墙
func GetPictures(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	imagesTypes, _ := strconv.Atoi(c.Query("type"))
	imagesHidden, _ := strconv.Atoi(c.Query("hidden"))
	imagesStatus, _ := strconv.Atoi(c.Query("status"))
	var pictures []models.PictureWall
	var count int
	var err error
	if page == 0 && limit == 0 {
		err = global.
			DB.
			Order("id desc").
			Preload("Author").
			Preload("Tags").
			Preload("ThumbsUpList").
			Preload("LeaveMessage").
			Preload("LeaveMessage.User").
			Where("type = ? and hidden = ? and status = ?", imagesTypes, imagesHidden, imagesStatus).
			Find(&pictures).
			Count(&count).
			Error
	} else {
		err = global.
			DB.
			Order("id desc").
			Limit(limit).
			Offset((page-1)*limit).
			Preload("Author").
			Preload("Tags").
			Preload("ThumbsUpList").
			Preload("LeaveMessage").
			Preload("LeaveMessage.User").
			Where("type = ? and hidden = ? and status = ?", imagesTypes, imagesHidden, imagesStatus).
			Find(&pictures).
			Limit(-1).
			Offset(-1).
			Count(&count).
			Error
	}
	if err == nil {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 1, Msg: "获取成功", Data: pictures, Count: count}),
		)
	} else {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 0, Msg: "获取失败", Error: err}),
		)
	}
}
