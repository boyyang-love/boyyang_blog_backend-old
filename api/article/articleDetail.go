/**
 * @Author: boyyang
 * @Date: 2022-07-11 12:51:43
 * @LastEditTime: 2022-07-11 13:02:32
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\article\articleDetail.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package api

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 根据id查询
func ArticleDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 0, Msg: "文章id不能为空"}),
		)
		return
	}
	var article models.Article
	global.
		DB.
		Preload("Author").
		Where("id = ?", id).
		Find((&article))
	c.JSON(
		http.StatusOK,
		utils.Msg(utils.Message{Code: 1, Msg: "获取成功", Data: article}),
	)
}
