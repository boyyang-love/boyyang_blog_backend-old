/**
 * @Author: boyyang
 * @Date: 2022-09-03 13:48:14
 * @LastEditTime: 2022-09-03 14:39:09
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog_server\api\article\editArticle.go
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

func EditArticle(c *gin.Context) {
	var article models.Article
	c.ShouldBind(&article)
	if article.ID != 0 {
		err := global.
			DB.
			Omit("Author").
			Model(&models.Article{}).
			Update(&article).
			Error
		if err == nil {
			c.JSON(http.StatusOK, utils.Msg(utils.Message{Code: 1, Msg: "修改成功"}))
		} else {
			c.JSON(http.StatusBadRequest, utils.Msg(utils.Message{Code: 0, Msg: "修改失败"}))
		}
	} else {
		c.JSON(http.StatusBadRequest, utils.Msg(utils.Message{Code: 0, Msg: "id 标题 描述 内容，为必填项"}))
	}
}
