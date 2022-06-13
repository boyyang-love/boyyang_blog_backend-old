/**
 * @Author: boyyang
 * @Date: 2022-04-10 21:31:13
 * @LastEditTime: 2022-06-13 13:04:01
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\user\updateUser.go
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

func UpdateUser(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, _ := utils.ParseToken(token)
	username := c.PostForm("username")
	qq := c.PostForm("qq")
	birthday, _ := strconv.ParseInt(c.PostForm("birthday"), 0, 64)
	sex, _ := strconv.Atoi(c.PostForm("sex"))
	blog_url := c.PostForm("blog_url")
	user := models.User{
		Username: username,
		Qq:       &qq,
		Birthday: &birthday,
		Sex:      &sex,
		BlogUrl:  &blog_url,
	}
	res := global.
		DB.
		Where("username = ? and id <> ?", username, claims.Id).
		Find(&models.User{})
	if res.RowsAffected > 0 {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "用户名已存在"}),
		)
	} else {
		err := global.
			DB.
			Model(models.User{}).
			Where("id = ?", claims.Id).
			Update(&user).
			Error
		if err == nil {
			c.JSON(
				http.StatusOK,
				utils.Msg(utils.Message{Code: 1, Msg: "更新成功"}),
			)
		} else {
			c.JSON(
				http.StatusBadRequest,
				utils.Msg(utils.Message{Code: 0, Msg: "更新失败"}),
			)
		}
	}
}
