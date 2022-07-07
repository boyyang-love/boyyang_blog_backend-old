/**
 * @Author: boyyang
 * @Date: 2022-06-03 11:11:28
 * @LastEditTime: 2022-07-07 19:40:30
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\loginRegister\register.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package api

import (
	"blog/global"
	"blog/models"
	"blog/setupSendEmail"
	"blog/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 注册
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	if strings.Trim(username, "") != "" &&
		strings.Trim(password, "") != "" &&
		strings.Trim(email, "") != "" {
		res := global.
			DB.
			Where("Username = ?", username).
			Or("Email = ?", email).
			Find(&models.User{})
		if res.RowsAffected == 0 {
			addUser := models.User{
				Username: username,
				Password: utils.MD5(password),
				Email:    &email,
			}
			err := global.
				DB.
				Create(&addUser).
				Error
			if err == nil {
				var pics []models.PictureWall
				urls := []string{}
				err = global.
					DB.
					Model(&models.PictureWall{}).
					Limit(20).
					Find(&pics).
					Limit(-1).
					Error
				if err == nil {
					for _, pic := range pics {
						urls = append(urls, pic.Url)
					}
				}
				setupSendEmail.SendEmail(
					[]string{email},
					"个人博客网站注册",
					setupSendEmail.EmailContent{
						Name: addUser.Username,
						Urls: urls,
					},
				)
				c.JSON(
					http.StatusOK,
					utils.Msg(utils.Message{Code: 1, Msg: "注册成功"}),
				)
			} else {
				c.JSON(
					http.StatusBadRequest,
					utils.Msg(utils.Message{Code: 0, Msg: "注册失败"}),
				)
			}
		} else {
			c.JSON(
				http.StatusBadRequest,
				utils.Msg(utils.Message{Code: 0, Msg: "该用户已存在"}),
			)
		}
	} else {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "账号,密码,邮箱为必填项"}),
		)
	}
}
