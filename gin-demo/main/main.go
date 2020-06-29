package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main()  {
	// Engin
	router := gin.Default()
	router.GET("/hello", hello) // hello函数处理"/hello"请求
	// 指定地址和端口号
	router.Run(":9090")
}
func hello(context *gin.Context) {
	println(">>>> hello function start <<<<")

	context.JSON(http.StatusOK,gin.H{
		"code":200,
		"success":true,
	})
}