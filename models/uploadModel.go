/*
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
	Url      string `json:"url"`
	Name     string `json:"file_name"`
	UploadID int    `json:"upload_id"`
}

func (Upload) TableName() string {
	return "uploads"
}
