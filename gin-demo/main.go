package main

import (
	"fmt"
	"gotrunks/gin-demo/config"
	"gotrunks/gin-demo/core/database"
	"gotrunks/gin-demo/core/log"
	"gotrunks/gin-demo/core/redis"
	"gotrunks/gin-demo/middleware"
	"gotrunks/gin-demo/router"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfig() //加载配置
}

func main() {
	r := gin.New()

	log.InitLog()          //配置日志
	redis.ConnectRedis()   //连接redis
	database.ConnectDB()   //连接数据库
	database.AutoMigrate() //自动迁移
	fmt.Print("进到这里了吗？\n")
	middleware.LoadMiddlewares(r) //加载中间件

	router.LoadRoutes(r)   //加载路由
	r.Run(config.App.Port) // listen and serve on 0.0.0.0:8080

	resourceRelease()
}

func resourceRelease() {
	go func() {
		database.DisconnectDB()
	}()
}
