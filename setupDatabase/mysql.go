/**
 * @Author: boyyang
 * @Date: 2022-02-14 10:43:46
 * @LastEditTime: 2022-06-10 08:35:21
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\setupDatabase\mysql.go
 */
package setupDatabase

import (
	"blog/global"
	"blog/setting"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//数据库连接
func SetupDB() *gorm.DB {
	driverName := "mysql"
	host := setting.MysqlHost //127.0.0.1
	port := setting.MysqlPort //3306
	database := setting.MysqlDatabase
	username := setting.MysqlUserName
	password := setting.MysqlPassword
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
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
