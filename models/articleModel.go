/*
 * @Author: boyyang
 * @Date: 2022-02-14 16:23:26
 * @LastEditTime: 2022-02-16 09:30:55
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\websit\models\articleModel.go
 */

package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Content  string `json:"content"`
	UserID   int    `json:"author_id"`
	Author   User   `json:"author" gorm:"foreignKey:UserID"`
}

func (Article) TableName() string {
	return "articles"
}
