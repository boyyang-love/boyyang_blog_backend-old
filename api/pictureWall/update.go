/**
 * @Author: boyyang
 * @Date: 2022-06-03 11:16:03
 * @LastEditTime: 2022-06-11 19:05:03
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\api\pictureWall\update.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package api

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 更新图片信息
func UpdatePicture(c *gin.Context) {
	tags := c.PostForm("tags")
	var form models.PictureWall
	var picture models.PictureWall
	c.ShouldBind(&form)
	// tags
	if strings.Trim(tags, " ") != "" {
		splitTags := strings.Split(tags, ",")
		for _, tag := range splitTags {
			form.Tags = append(form.Tags, models.ImagesTag{
				TagName: tag,
			})
		}
	}
	err := global.
		DB.
		Debug().
		Omit("Author").
		Model(&picture).
		Update(&form).
		Association("Tags").
		Replace(form.Tags).
		Error
	if err == nil {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 1, Msg: "更新成功", Data: picture}),
		)
	} else {
		c.JSON(
			http.StatusBadRequest,
			utils.Message{Code: 0, Msg: "更新失败", Error: err},
		)
	}
}
