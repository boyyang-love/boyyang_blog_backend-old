/**
 * @Author: boyyang
 * @Date: 2022-02-14 16:23:26
 * @LastEditTime: 2022-02-19 10:34:08
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\models\articleModel.go
 */

package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title    string `json:"title"`                           //文章标题
	Subtitle string `json:"subtitle"`                        //文章副标题
	Image    string `json:"image"`                           //文章图片
	Content  string `json:"content"`                         //文章内容
	ThumbsUp int    `json:"thumbs_up" grom:"default 0"`      //点赞数
	UserID   int    `json:"author_id"`                       //作者id
	Author   User   `json:"author" gorm:"foreignKey:UserID"` //作者信息
}

func (Article) TableName() string {
	return "articles"
}
