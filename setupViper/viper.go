/**
 * @Author: boyyang
 * @Date: 2022-06-10 17:39:26
 * @LastEditTime: 2022-06-10 18:59:44
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\setupViper\viper.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package setupViper

import (
	"blog/global"
	"fmt"

	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("config.yml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("💔viper读取配置文件失败,请检查配置文件是否存在!", err)
	} else {
		fmt.Println("💞🎈🎐viper读取配置文件成功!")
	}
	if err = v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}
	global.Viper = v
	return v
}
