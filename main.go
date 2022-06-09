/**
 * @Author: boyyang
 * @Date: 2022-02-13 19:45:12
 * @LastEditTime: 2022-06-09 11:21:17
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\main.go
 */

package main

import (
	"blog/global"
	"blog/setupClient"
	setupDatabase "blog/setupDatabase"
	setupRouter "blog/setupRouter"
)

func main() {
	// mysql数据库初始化
	setupDatabase.SetupDB()
	if global.DB != nil {
		setupDatabase.RegisterTables(global.DB) // 初始化表
		// 程序结束前关闭数据库链接
		db := global.DB.DB()
		defer db.Close()
	}
	setupClient.SetupClient()
	// 路由初始化
	setupRouter.SetupRouter()
}
