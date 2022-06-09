/**
 * @Author: boyyang
 * @Date: 2022-06-09 09:10:01
 * @LastEditTime: 2022-06-09 13:14:28
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\global\global.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */
package global

import (
	"github.com/jinzhu/gorm"
	"github.com/tencentyun/cos-go-sdk-v5"
)

var DB *gorm.DB        // 数据库实例
var Client *cos.Client // cos实例
