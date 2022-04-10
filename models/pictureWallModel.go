/**
 * @Author: boyyang
 * @Date: 2022-04-03 00:02:39
 * @LastEditTime: 2022-04-10 18:03:26
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\models\pictureWallModel.go
 * [如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package models

import "github.com/jinzhu/gorm"

type PictureWall struct {
	gorm.Model
	Url    string `json:"url" form:"url"`                        //图片地址
	Name   string `json:"name" form:"name"`                      //图片名称
	Des    string `json:"des" form:"des"`                        //图片描述
	Type   int    `json:"type" form:"type" gorm:"default:0"`     //图片类型 0 pc端 1 手机端
	Hidden int    `json:"hidden" form:"hidden" gorm:"default:0"` //是否隐藏 0 否 1是
	UserID int    `json:"user_id"`                               //用户id
	Author User   `json:"author,omitempty" gorm:"foreignKey:UserID"`
}

func (PictureWall) TableName() string {
	return "pictureWalls"
}
