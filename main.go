/**
 * @Author: boyyang
 * @Date: 2022-02-13 19:45:12
 * @LastEditTime: 2022-02-14 10:51:43
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \go-study\websit\main.go
 */

package main

import (
	setupDatabase "websit/setupDatabase"
	setupRouter "websit/setupRouter"
)

func main() {
	// mysql数据库初始化
	db := setupDatabase.SetupDB()
	defer db.Close() //延时关闭

	r := setupRouter.SetupRouter()

	r.Run(":80")
}
