/**
 * @Author: boyyang
 * @Date: 2022-07-11 12:53:42
 * @LastEditTime: 2022-07-11 12:53:44
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\article\delArticle.go
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

// 删除文章
func DelArticle(c *gin.Context) {
	id := c.PostForm("id")
	var articel models.Article
	err := global.
		DB.Where("id = ?", id).
		Delete(&articel).
		Error
	if err == nil {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 1, Msg: "删除成功"}),
		)
	} else {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "删除失败", Error: err}),
		)
	}
}
