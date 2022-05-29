/**
 * @Author: boyyang
 * @Date: 2022-02-13 19:45:12
 * @LastEditTime: 2022-05-28 20:05:14
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
