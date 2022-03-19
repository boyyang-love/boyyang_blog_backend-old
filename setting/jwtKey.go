/**
 * @Author: boyyang
 * @Date: 2022-02-18 14:18:44
 * @LastEditTime: 2022-02-18 15:19:53
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\setting\jwtKey.go
 */

package setting

var JwtSecret []byte

func init() {
	JwtSecret = []byte("YZH")
}
