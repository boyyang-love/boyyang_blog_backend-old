/**
 * @Author: boyyang
 * @Date: 2022-06-24 09:49:37
 * @LastEditTime: 2022-06-28 13:59:00
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\models\thumbsUp.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */
package models

import "time"

type ThumbsUp struct {
	// gorm.Model
	ID         uint       `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index" json:"deleted_at"`
	UserId     int        `json:"user_id" form:"user_id"`
	ThumbsUpId int        `json:"thumbs_up_id" form:"thumbs_up_id"`
	Author     User       `json:"author" gorm:"foreignKey:UserId"`
}

func (ThumbsUp) TableName() string {
	return "thumbsUp"
}
