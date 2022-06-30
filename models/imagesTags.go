/**
 * @Author: boyyang
 * @Date: 2022-04-23 17:00:17
 * @LastEditTime: 2022-06-30 08:50:24
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\models\imagesTags.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package models

import "time"

type ImagesTag struct {
	// gorm.Model
	ID            uint       `gorm:"primary_key" json:"id" form:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `sql:"index" json:"deleted_at"`
	TagName       string     `json:"tag_name"`        //标签名称
	PictureWallID int        `json:"picture_wall_id"` //用户id
}

func (ImagesTag) TableName() string {
	return "imagesTag"
}
