/**
 * @Author: boyyang
 * @Date: 2022-02-18 14:18:44
 * @LastEditTime: 2022-06-10 17:00:55
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

// 数据库配置
var (
	MysqlHost     string = "sh-cynosdbmysql-grp-q87hvhhs.sql.tencentcdb.com" //数据库地址
	MysqlPort     string = "26019"                                           //数据库端口
	MysqlUserName string = "root"                                            //数据库用户名
	MysqlPassword string = "zxgf8bTa"                                        //数据库密码
	MysqlDatabase string = "boyyang"                                         //数据库名称
)

// 腾讯云 对象存储配置
var (
	ClientUrl       string = "https://7072-prod-2g489qm8208c3cfd-1301921121.cos.ap-shanghai.myqcloud.com" //对象存储地址
	ClientSecretID  string = "AKID85lAJBm6AyVljQr0sSmpfKn0vZw3sQhh"                                       // id
	ClientSecretKey string = "kDPLzNUKtBk23nGRM0MREkzxd3C4f03p"                                           // key
)
