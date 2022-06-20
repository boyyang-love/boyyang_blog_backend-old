/**
 * @Author: boyyang
 * @Date: 2022-06-10 17:56:21
 * @LastEditTime: 2022-06-14 10:15:03
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\config\config.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package config

type Server struct {
	Mysql     Mysql     `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	CloudBase CloudBase `mapstructure:"cloudBase" json:"cloud_base" yaml:"cloudBase"`
	Email     Email     `mapstructure:"email" json:"email" yaml:"email"`
	Servers   Servers   `mapstructure:"servers" json:"servers" yaml:"servers"`
}
