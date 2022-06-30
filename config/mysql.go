/**
 * @Author: boyyang
 * @Date: 2022-06-10 17:02:27
 * @LastEditTime: 2022-06-30 08:33:56
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\config\mysql.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package config

type Mysql struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
	Timeout  string `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
}
