/**
 * @Author: boyyang
 * @Date: 2022-06-03 11:14:29
 * @LastEditTime: 2022-07-18 14:44:27
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\pictureWall\uploadPicture.go
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

// 上传图片信息
func UploadPicture(c *gin.Context) {
	token := c.Request.Header.Get("token")
	tags := c.PostForm("tags")
	claims, _ := utils.ParseToken(token)
	var form models.PictureWall
	c.ShouldBind(&form)
	form.UserID = claims.Id
	// tags
	if strings.Trim(tags, "") != "" {
		splitTags := strings.Split(tags, ",")
		for _, tag := range splitTags {
			form.Tags = append(form.Tags, models.ImagesTag{
				TagName: tag,
			})
		}
	}
	res := global.
		DB.
		Create(&form)
	if res.Error == nil {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 1, Msg: "作品上传成功", Data: form}),
		)
	} else {
		c.JSON(
			http.StatusBadRequest,
			utils.Message{Code: 0, Msg: "作品上传失败", Error: res.Error},
		)
	}
}
