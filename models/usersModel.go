/*
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
	Username  string `json:"username"`
	Password  string `json:"password"`
	AvaterUrl string `json:"avaterUrl"`
	Age       int    `json:"age"`
	Sex       string `json:"sex"`
}

//表示把 User 结构体默认操作的表改为 user 表
func (User) TableName() string {
	return "users"
}
