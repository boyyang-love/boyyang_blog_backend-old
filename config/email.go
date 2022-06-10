/**
 * @Author: boyyang
 * @Date: 2022-06-10 19:28:56
 * @LastEditTime: 2022-06-10 19:30:58
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\config\email.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package config

type Email struct {
	From    string `mapstructure:"from" json:"from" yaml:"from"`
	SmtpKey string `mapstructure:"smtpKey" json:"smtp_key" yaml:"smtpKey"`
}
