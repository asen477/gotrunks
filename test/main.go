package main

import (
	"fmt"
	"hello/config"
)

func init() {
	config.LoadConfig() //加载配置
}

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(config.Name)
}