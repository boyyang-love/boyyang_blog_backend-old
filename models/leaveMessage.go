/**
 * @Author: boyyang
 * @Date: 2022-07-13 18:14:48
 * @LastEditTime: 2022-07-13 18:24:35
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\models\leaveMessage.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package models

import "time"

type LeaveMessage struct {
	ID        uint       `gorm:"primary_key" json:"id" form:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	CommentId int        `json:"comment_id"` //评论id
	Message   string     `json:"message"`    //留言内容
	UserID    int        `json:"user_id"`    //用户id
	User      User       `json:"user" gorm:"fareignKey:UserID"`
}

func (LeaveMessage) TableName() string {
	return "leaveMessages"
}
