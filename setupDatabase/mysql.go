/**
 * @Author: boyyang
 * @Date: 2022-02-14 10:43:46
 * @LastEditTime: 2022-05-21 15:10:53
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\setupDatabase\mysql.go
 */
package setupDatabase

import (
	"blog/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

//数据库连接
func SetupDB() *gorm.DB {
	driverName := "mysql"
	host := "sh-cynosdbmysql-grp-q87hvhhs.sql.tencentcdb.com" //127.0.0.1
	port := "26019"                                           //3306
	database := "boyyang"
	username := "root"
	password := "zxgf8bTa" // root
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
		fmt.Println(args)
		fmt.Println("<<<<mysql连接成功>>>>")
	}
	DB.AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Upload{},
		&models.PictureWall{},
		&models.ImagesTag{},
	)
	// DB.AutoMigrate(&models.Article{})
	// DB.AutoMigrate(&models.Upload{})
	return DB
}
