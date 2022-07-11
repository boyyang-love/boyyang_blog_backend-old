/**
 * @Author: boyyang
 * @Date: 2022-07-11 12:50:25
 * @LastEditTime: 2022-07-11 13:03:02
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\article\getArticles.go
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

// 获取所有文章
func Articles(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")
	_page, _ := strconv.ParseInt(page, 10, 32)
	_limit, _ := strconv.ParseInt(limit, 10, 32)
	var articles []models.Article
	var count int
	if _page == 0 && _limit == 0 {
		global.
			DB.
			Preload("Author").
			Find(&articles).
			Count(&count)
	} else {
		global.
			DB.
			Limit(_limit).
			Offset((_page - 1) * _limit).
			Preload("Author").
			Find(&articles).
			Offset(-1).
			Limit(-1).
			Count(&count)
	}
	c.JSON(
		http.StatusOK,
		utils.Msg(utils.Message{Code: 1, Msg: "获取成功", Data: articles, Count: count}),
	)
}
