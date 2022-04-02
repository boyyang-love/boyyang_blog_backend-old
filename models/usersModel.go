/**
 * @Author: boyyang
 * @Date: 2022-02-14 11:00:21
 * @LastEditTime: 2022-02-18 17:30:48
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\models\usersModel.go
 */

package models

import (
	"github.com/jinzhu/gorm"
)

//定义 user 模型：
type User struct { // 默认表名是 `users`
	gorm.Model
	Username   string `json:"username"`                   //用户名
	Password   string `json:"password"`                   //密码
	AvaterUrl  string `json:"avaterUrl"`                  //头像地址
	Background string `json:"background_img"`             //背景图片
	Age        int    `json:"age"`                        //年龄
	Sex        int    `json:"sex" gorm:"default: 1"`      //性别 0 女 1 男
	Birthday   int    `json:"birthday"`                   //生日
	Role       string `json:"role" gorm:"default:'user'"` //角色
	TelPhone   string `json:"tel"`                        //手机号
	Email      string `json:"email"`                      //email
}

//表示把 User 结构体默认操作的表改为 user 表
func (User) TableName() string {
	return "users"
}
