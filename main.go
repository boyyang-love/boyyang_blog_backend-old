/**
 * @Author: boyyang
 * @Date: 2022-02-13 19:45:12
 * @LastEditTime: 2022-06-09 16:25:18
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
		// 初始化表
		setupDatabase.RegisterTables(global.DB)
		db := global.DB.DB()
		defer db.Close()
	}
	// 对象存储初始化
	setupClient.SetupClient()
	// 路由初始化
	setupRouter.SetupRouter()
}
