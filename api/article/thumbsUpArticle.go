/**
 * @Author: boyyang
 * @Date: 2022-02-16 10:20:47
 * @LastEditTime: 2022-07-11 13:03:41
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\article\thumbsUpArticle.go
 * [如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package api

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// 文章点赞
func ThumbsUpArticle(c *gin.Context) {
	id := c.Query("id")
	err := global.
		DB.
		Model(&models.Article{}).
		Where("id = ?", id).
		Update("thumbs_up", gorm.Expr("thumbs_up + 1")).
		Error
	if err == nil {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 1, Msg: "点赞成功"}),
		)
	}
}
