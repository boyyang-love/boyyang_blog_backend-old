/**
 * @Author: boyyang
 * @Date: 2022-07-11 12:52:51
 * @LastEditTime: 2022-07-11 12:52:53
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\article\addArticle.go
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

// 添加文章
func AddArticle(c *gin.Context) {
	title := c.PostForm("title")
	subtitle := c.PostForm("subtitle")
	content := c.PostForm("content")
	image := c.PostForm("image")
	token := c.Request.Header["Token"]
	var author_id int
	userInfo, err := utils.ParseToken(token[0])
	if err == nil {
		author_id = userInfo.Id
	}
	if len(title) == 0 {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "文章标题不能为空"}),
		)
	} else {
		article := models.Article{
			Title:    title,
			Subtitle: subtitle,
			Image:    image,
			Content:  content,
			UserID:   author_id,
		}
		err := global.
			DB.
			Create(&article).
			Error
		if err == nil {
			c.JSON(
				http.StatusOK,
				utils.Msg(utils.Message{Code: 1, Msg: "文章添加成功", Data: article.ID}),
			)
		} else {
			c.JSON(
				http.StatusBadRequest,
				utils.Msg(utils.Message{Code: 0, Msg: "文章添加失败"}),
			)
		}
	}
}
