/**
 * @Author: boyyang
 * @Date: 2022-02-19 10:48:12
 * @LastEditTime: 2022-04-23 16:52:12
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\models\uploads.go
 */

package models

import "github.com/jinzhu/gorm"

type Upload struct {
	gorm.Model
	Url      string `json:"url"`                             //图片地址
	FileName string `json:"file_name"`                       //图片名称
	UserID   int    `json:"user_id"`                         //用户id
	Author   User   `json:"author" gorm:"foreignKey:UserID"` //用户
}

func (Upload) TableName() string {
	return "uploads"
}
