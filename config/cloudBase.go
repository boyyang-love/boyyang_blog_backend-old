/**
 * @Author: boyyang
 * @Date: 2022-06-10 18:02:11
 * @LastEditTime: 2022-06-10 18:54:20
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\config\cloudBase.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package config

type CloudBase struct {
	ClientUrl       string `mapstructure:"clientUrl" json:"client_url" yaml:"clientUrl"`
	ClientSecretID  string `mapstructure:"clientSecretId" json:"client_secret_id" yaml:"clientSecretId"`
	ClientSecretKey string `mapstructure:"clientSecretKey" json:"client_secret_key" yaml:"clientSecretKey"`
}
