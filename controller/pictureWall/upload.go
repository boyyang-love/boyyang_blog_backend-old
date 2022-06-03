/**
 * @Author: boyyang
 * @Date: 2022-06-03 11:14:29
 * @LastEditTime: 2022-06-03 11:14:30
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\controller\pictureWall\upload.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */
package controller

import (
	"blog/models"
	"blog/setupDatabase"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 上传图片信息
func UploadPicture(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, _ := utils.ParseToken(token)
	var form models.PictureWall
	c.Bind(&form)
	form.UserID = claims.Id
	res := setupDatabase.
		DB.
		Create(&form)
	if res.Error == nil {
		c.JSON(
			http.StatusOK,
			utils.Msg(utils.Message{Code: 1, Msg: "上传成功", Data: form.ID}),
		)
	} else {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "上传失败", Data: res.Error}),
		)
	}
}
