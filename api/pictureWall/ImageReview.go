/**
 * @Author: boyyang
 * @Date: 2022-07-19 13:46:43
 * @LastEditTime: 2022-07-20 14:56:33
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\pictureWall\ImageReview.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package api

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 图片状态修改
func ImageReview(c *gin.Context) {
	id := c.Query("id")
	_type := c.Query("type")
	status, _ := strconv.Atoi(c.Query("status"))
	if strings.Trim(id, "") != "" {
		if _type == "type" || _type == "status" || _type == "hidden" {
			err := global.
				DB.
				Debug().
				Model(&models.PictureWall{}).
				Where("id = ?", id).
				Update(_type, status).
				Error
			if err == nil {
				c.JSON(
					http.StatusOK,
					utils.Msg(utils.Message{Code: 1, Msg: "更新状态成功"}),
				)
			} else {
				c.JSON(
					http.StatusBadRequest,
					utils.Msg(utils.Message{Code: 0, Msg: "更新失败", Error: err}),
				)
			}
		} else {
			c.JSON(
				http.StatusBadRequest,
				utils.Msg(utils.Message{Code: 0, Msg: "参数错误"}),
			)
		}
	} else {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "id不能为空"}),
		)
	}
}
