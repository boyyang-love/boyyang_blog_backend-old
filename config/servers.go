/**
 * @Author: boyyang
 * @Date: 2022-06-14 10:14:11
 * @LastEditTime: 2022-06-14 10:20:39
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\config\servers.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package config

type Servers struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
}
