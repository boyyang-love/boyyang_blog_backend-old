/**
 * @Author: boyyang
 * @Date: 2022-04-03 00:35:57
 * @LastEditTime: 2022-06-09 09:15:15
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\controller\pictureWall\pictureDetail.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package controller

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取图片详情
func GetPictureDetail(c *gin.Context) {
	id := c.Param("id")
	var picture models.PictureWall
	err := global.
		DB.
		First(&picture, id).
		Error
	if err == nil {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 1, Msg: "获取成功", Data: picture}),
		)
	} else {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 0, Msg: "获取失败", Error: err}),
		)
	}
}
