/**
 * @Author: boyyang
 * @Date: 2022-02-19 10:48:12
 * @LastEditTime: 2022-06-28 13:59:11
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\models\uploads.go
 */

package models

import "time"

type Upload struct {
	// gorm.Model
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Url       string     `json:"url"`                             //图片地址
	FileName  string     `json:"file_name"`                       //图片名称
	UserID    int        `json:"user_id"`                         //用户id
	Author    User       `json:"author" gorm:"foreignKey:UserID"` //用户
}

func (Upload) TableName() string {
	return "uploads"
}
