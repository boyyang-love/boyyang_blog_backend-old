/**
 * @Author: boyyang
 * @Date: 2022-06-03 11:19:09
 * @LastEditTime: 2022-07-15 18:45:24
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\pictureWall\deletePicture.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package api

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 删除图片
func DeletePicture(c *gin.Context) {
	id := c.Query("id")
	var picture models.PictureWall
	if strings.Trim(id, "") == "" {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "id不能为空"}),
		)
	} else {
		err := global.
			DB.
			Debug().
			Where("id = ?", id).
			First(&picture).
			Delete(&picture).
			Error
		if err == nil {
			path := picture.Url
			_, err = global.Client.Object.Delete(context.Background(), path)
			if err == nil {
				c.JSON(
					http.StatusOK,
					utils.Msg(utils.Message{Code: 1, Msg: "删除成功"}),
				)
			} else {
				c.JSON(
					http.StatusBadRequest,
					utils.Msg(utils.Message{Code: 1, Msg: "对象存储图片删除失败", Error: err}),
				)
			}
		} else {
			c.JSON(
				http.StatusOK,
				utils.Msg(utils.Message{Code: 0, Msg: "删除失败", Error: err}),
			)
		}
	}
}
