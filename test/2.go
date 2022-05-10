package main

import (
	"fmt"
	"hello/config"
)

var age = test()

// 为了看到全局变量是先被初始化,先写个函数
func test() int {
	fmt.Println("test()") // 1
	return 90
}

// init 函数,通常可以在 init 函数中完成初始化工作
func init() {
	fmt.Println("init()...") // 2
}

func main() {
	fmt.Println("Age=", config.Age, "Name=", config.Name)
	fmt.Println("main()...")
}