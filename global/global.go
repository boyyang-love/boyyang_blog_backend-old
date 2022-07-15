/**
 * @Author: boyyang
 * @Date: 2022-06-09 09:10:01
 * @LastEditTime: 2022-07-15 18:18:54
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\global\global.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */
package global

import (
	"blog/config"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
)

var (
	DB     *gorm.DB      // 数据库实例
	Client *cos.Client   // cos实例
	Config config.Config // 配置
	Viper  *viper.Viper  // viper实例
)
