/**
 * @Author: boyyang
 * @Date: 2022-04-23 17:00:17
 * @LastEditTime: 2022-05-21 20:44:30
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\models\imagesTags.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package models

import "github.com/jinzhu/gorm"

type ImagesTag struct {
	gorm.Model
	TagName       string `json:"tag_name"` //标签名称
	PictureWallID int    `json:"picturewall_id"`
}

func (ImagesTag) TableName() string {
	return "imagesTag"
}
