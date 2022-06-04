/**
 * @Author: boyyang
 * @Date: 2022-06-03 11:19:09
 * @LastEditTime: 2022-06-04 11:23:38
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\controller\pictureWall\deletePicture.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package controller

import (
	"blog/models"
	"blog/setupDatabase"
	"blog/utils"
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
		err := setupDatabase.
			DB.
			Debug().
			Where("id = ?", id).
			Delete(&picture).
			Error
		if err == nil {
			c.JSON(
				http.StatusOK,
				utils.Msg(utils.Message{Code: 1, Msg: "删除成功"}),
			)
		} else {
			c.JSON(
				http.StatusOK,
				utils.Msg(utils.Message{Code: 0, Msg: "删除失败", Error: err}),
			)
		}
	}

}
