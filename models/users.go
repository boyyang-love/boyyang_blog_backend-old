/**
 * @Author: boyyang
 * @Date: 2022-02-14 11:00:21
 * @LastEditTime: 2022-06-30 08:50:52
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\models\users.go
 */

package models

import "time"

//定义 user 模型：
type User struct { // 默认表名是 `users`
	// gorm.Model
	ID            uint       `gorm:"primary_key" json:"id" form:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `sql:"index" json:"deleted_at"`
	Username      string     `json:"username" form:"username"`               //用户名
	Password      string     `json:"password"`                               //密码
	AvaterUrl     string     `json:"avater_url" form:"avater_url"`           //头像地址
	BackgroundImg string     `json:"background_img" form:"background_img"`   //背景图片地址
	Age           int        `json:"age" form:"age"`                         //年龄
	Sex           *int       `json:"sex" gorm:"default: 1" form:"sex"`       //性别 0 女 1 男
	Birthday      *int64     `json:"birthday" form:"birthday"`               //生日
	Role          string     `json:"role" gorm:"default:'user'" form:"role"` //角色
	TelPhone      *string    `json:"tel_phone" form:"tel_phone"`             //手机号
	Email         *string    `json:"email" form:"email"`                     //email
	Qq            *string    `json:"qq" form:"qq"`                           //qq
	Wechat        string     `json:"wechat" form:"wechat"`                   //微信
	BlogUrl       *string    `json:"blog_url" form:"blog_url"`               //博客地址
	Address       string     `json:"address" form:"address"`                 //地址
}

//表示把 User 结构体默认操作的表改为 user 表
func (User) TableName() string {
	return "users"
}
