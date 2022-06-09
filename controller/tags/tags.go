/**
 * @Author: boyyang
 * @Date: 2022-04-29 11:18:05
 * @LastEditTime: 2022-06-09 09:15:44
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\controller\tags\tags.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package controllers

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 添加tags
func AddTags(c *gin.Context) {
	tag_name := c.PostForm("tag_name")
	tags := models.ImagesTag{
		TagName: tag_name,
	}
	if tag_name != "" {
		res := global.
			DB.
			Where("tag_name = ?", tag_name).
			Find(&tags)
		if res.RowsAffected == 0 {
			global.
				DB.
				Create(&tags)
			c.JSON(
				http.StatusOK,
				utils.Msg(utils.Message{Code: 1, Msg: "添加成功"}),
			)
		} else {
			c.JSON(
				http.StatusBadRequest,
				utils.Msg(utils.Message{Code: 0, Msg: "标签已存在"}),
			)
		}
	} else {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "标签名不能为空"}),
		)
	}
}
