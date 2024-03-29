/**
 * @Author: boyyang
 * @Date: 2022-02-14 16:23:26
 * @LastEditTime: 2022-09-03 14:38:15
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog_server\models\articles.go
 */

package models

import "time"

type Article struct {
	// gorm.Model
	ID        uint       `gorm:"primary_key" json:"id" form:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Title     string     `json:"title" form:"title"`                       //文章标题
	Subtitle  string     `json:"subtitle" form:"subtitle"`                 //文章副标题
	Image     string     `json:"image"`                                    //文章图片
	Content   string     `json:"content" form:"content" gorm:"size:15000"` //文章内容
	ThumbsUp  int        `json:"thumbs_up" grom:"default 0"`               //点赞数
	UserID    int        `json:"author_id"`                                //作者id
	Author    User       `json:"author" gorm:"foreignKey:UserID"`          //作者信息
}

func (Article) TableName() string {
	return "articles"
}
