/**
 * @Author: boyyang
 * @Date: 2022-02-18 14:18:44
 * @LastEditTime: 2022-06-11 12:43:03
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\setting\setting.go
 */

package setting

// JWT配置
var (
	JwtSecret []byte = []byte("YZH") //JWT密钥
)

// 服务器配置
var (
	ServerPort string = ":80" //服务器端口
)
