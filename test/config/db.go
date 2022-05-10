package config

import (
	"fmt"
)

func LoadConfig() {
	fmt.Print("进到这里了吗？")
	fmt.Print("进到这里了吗？")
}

// init 函数,通常可以在 init 函数中完成初始化工作
func init() {
	fmt.Println("init()...") //2
}

func main() {
	fmt.Println("main()...")
}