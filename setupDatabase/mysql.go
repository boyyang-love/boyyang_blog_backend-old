/**
 * @Author: boyyang
 * @Date: 2022-02-14 10:43:46
 * @LastEditTime: 2022-06-30 13:06:12
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\setupDatabase\mysql.go
 */
package setupDatabase

import (
	"blog/global"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//数据库连接
func SetupDB() *gorm.DB {
	mysqlConfig := global.Config.Mysql
	driverName := "mysql"
	host := mysqlConfig.Host
	port := mysqlConfig.Port
	database := mysqlConfig.Database
	username := mysqlConfig.Username
	password := mysqlConfig.Password
	timeout := mysqlConfig.Timeout
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&writeTimeout=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		timeout,
	)
	var err error
	global.DB, err = gorm.Open(driverName, args)
	if err != nil {
		panic("💔数据库连接失败" + err.Error())
	} else {
		fmt.Println("💞🎈🎐数据库初始化成功")
	}
	return global.DB
}
