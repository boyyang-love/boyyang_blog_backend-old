/**
 * @Author: boyyang
 * @Date: 2022-02-13 19:45:12
 * @LastEditTime: 2022-04-05 15:36:31
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\main.go
 */

package main

import (
	setupDatabase "blog/setupDatabase"
	setupRouter "blog/setupRouter"
)

func main() {
	// mysql数据库初始化
	db := setupDatabase.SetupDB()
	defer db.Close() //延时关闭

	r := setupRouter.SetupRouter()

	r.Run(":80")
}
