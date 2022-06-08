/**
 * @Author: boyyang
 * @Date: 2022-02-13 19:45:12
 * @LastEditTime: 2022-06-08 13:51:47
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\main.go
 */

package main

import (
	setupDatabase "blog/setupDatabase"
	setupRouter "blog/setupRouter"
	"log"
)

func main() {
	// mysql数据库初始化
	db := setupDatabase.SetupDB()
	defer db.Close() //延时关闭
	err := db.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}
	r := setupRouter.SetupRouter()
	r.Run(":80")
}
