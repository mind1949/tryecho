package main

import (
	"iserver/controllers"
	"iserver/db"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	defer db.MysqlDB.Close()
	// 实例化一个Echo类型
	e := echo.New()
	// 添加中间件
	e.Use(middleware.Logger())
	// 添加路由
	e.GET("users/:id", controllers.FindUser, validateSession)
	// 启动服务器
	e.Logger.Fatal(e.Start(":1323"))
}
