/**
 * @Author: boyyang
 * @Date: 2022-06-22 13:35:31
 * @LastEditTime: 2022-06-24 10:16:10
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\pictureWall\thumbsUpPicture.go
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

func ThumbsUpPicture(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	token := c.Request.Header.Get("token")
	claims, _ := utils.ParseToken(token)
	var pictureWall models.PictureWall
	var thumbsUp models.ThumbsUp
	res := global.
		DB.
		Model(&thumbsUp).
		Where("user_id = ? AND thumbs_up_id = ?", claims.Id, id).
		First(&thumbsUp)
	if res.RowsAffected != 0 {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "已经点过赞了"}),
		)
		return
	}
	err := global.
		DB.
		Model(&pictureWall).
		Where("id = ?", id).
		First(&pictureWall).
		Error
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "不存在该图片", Error: err}),
		)
	} else {
		err := global.
			DB.
			Create(&models.ThumbsUp{UserId: claims.Id, ThumbsUpId: id}).
			Error
		if err == nil {
			pictureWall.ThumbsUp += 1
			global.
				DB.
				Save(&pictureWall)
			c.JSON(
				http.StatusOK,
				utils.Msg(utils.Message{Code: 1, Msg: "点赞成功"}),
			)
		} else {
			c.JSON(
				http.StatusBadRequest,
				utils.Msg(utils.Message{Code: 0, Msg: "点赞失败", Error: err}),
			)
		}
	}
}
