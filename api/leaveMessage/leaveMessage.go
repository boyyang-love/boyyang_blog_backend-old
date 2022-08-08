/**
 * @Author: boyyang
 * @Date: 2022-07-14 12:20:24
 * @LastEditTime: 2022-08-04 13:00:47
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\api\leaveMessage\leaveMessage.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package api

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func LeaveMessageTo(c *gin.Context) {
	token := c.Request.Header.Get("token")
	message := c.PostForm("message")
	commentId, _ := strconv.Atoi(c.PostForm("comment_id"))
	claims, _ := utils.ParseToken(token)
	if commentId == 0 {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "comment_id不能为空"}),
		)
	}
	if strings.Trim(message, "") != "" {
		addMessage := models.LeaveMessage{
			Message:   message,
			UserID:    claims.Id,
			CommentId: int(commentId),
		}
		err := global.DB.Create(&addMessage).Error
		if err == nil {
			c.JSON(
				http.StatusOK,
				utils.Msg(utils.Message{Code: 1, Msg: "留言成功"}),
			)
		} else {
			c.JSON(
				http.StatusBadRequest,
				utils.Msg(utils.Message{Code: 0, Msg: "留言失败"}),
			)
		}
	} else {
		c.JSON(
			http.StatusBadRequest,
			utils.Msg(utils.Message{Code: 0, Msg: "留言内容不能为空"}),
		)
	}
}
