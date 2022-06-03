/**
 * @Author: boyyang
 * @Date: 2022-06-03 11:16:03
 * @LastEditTime: 2022-06-03 11:16:05
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\controller\pictureWall\update.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package controller

import (
	"blog/models"
	"blog/setupDatabase"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 更新图片信息
func UpdatePicture(c *gin.Context) {
	var form models.PictureWall
	var picture models.PictureWall
	c.Bind(&form)
	err := setupDatabase.
		DB.
		Debug().
		Omit("Author").
		Model(&picture).
		Update(&form).
		Error
	if err == nil {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 1, Msg: "更新成功", Data: picture}),
		)
	} else {
		c.JSON(
			http.StatusBadRequest,
			utils.Message{Code: 0, Msg: "更新失败", Error: err},
		)
	}
}
