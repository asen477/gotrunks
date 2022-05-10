package config

import (
	"fmt"
)

var Age int
var Name string

// Age 和 Name 全局变量，我们需要在main.go 使用
// 但是我们需要初始化 Age 和 Name
// init 函数完成初始化工作
func init() {
	fmt.Println("utils 包的  init()...")
	Age = 100
	Name = "你好我是杨荔婷"
}
