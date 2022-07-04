/**
 * @Author: boyyang
 * @Date: 2022-07-03 12:02:04
 * @LastEditTime: 2022-07-03 13:39:26
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\models\loggers.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package models

import "time"

type Logger struct {
	ID            uint       `gorm:"primary_key" json:"id" form:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `sql:"index" json:"deleted_at"`
	Uid           uint       `json:"uid" gorm:"default:0"`
	Uri           string     `json:"uri"`
	Status        int        `json:"status"`
	Host          string     `json:"host"`
	ClientIP      string     `json:"client_ip"`
	RawQuery      string     `json:"raw_query"`
	StartTime     int64      `json:"start_time"`
	EndTime       int64      `json:"end_time"`
	RequestMethod string     `json:"request_method"`
	ContentType   string     `json:"content_type"`
	UserAgent     string     `json:"user_agent"`
}

func (Logger) TableName() string {
	return "loggers"
}
