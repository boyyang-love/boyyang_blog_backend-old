/*
 * @Author: boyyang
 * @Date: 2022-02-14 10:43:46
 * @LastEditTime: 2022-02-18 14:47:44
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\src\setupDatabase\mysql.go
 */
package setupDatabase

import (
	"fmt"
	"websit/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

//数据库连接
func SetupDB() *gorm.DB {
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "boyyang"
	username := "root"
	password := "root"
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

	DB, err = gorm.Open(driverName, args)
	//db, err := gorm.Open("mysql", "user:islot@/blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	} else {
		fmt.Println("<<<<mysql连接成功>>>>")
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Article{})

	return DB
}
