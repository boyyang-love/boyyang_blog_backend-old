/**
 * @Author: boyyang
 * @Date: 2022-02-19 10:48:12
 * @LastEditTime: 2022-02-19 11:40:31
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\models\uploadModel.go
 */

package models

import "github.com/jinzhu/gorm"

type Upload struct {
	gorm.Model
	Url    string `json:"url"`       //图片地址
	Name   string `json:"file_name"` //图片名称
	UserID int    `json:"user_id"`   //用户id
	Author User   `json:"author" gorm:"foreignKey:UserID"`
}

func (Upload) TableName() string {
	return "uploads"
}
