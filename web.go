package main

import (
	"fmt"
	"net/http"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是一个开始")
}

func main() {
	http.HandleFunc("/", myWeb)

	//fmt.Println("服务器即将开启，访问地址 http://140.143.205.57:9999")
	fmt.Println("服务器即将开启，访问地址 http://127.0.0.1:8090")

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("服务器开启错误: ", err)
	}
}
