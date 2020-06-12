/**
 * User: Trunks(Gao)
 * Date: 2020/06/10
 * goroutine协程
 */

package main

import (
	"fmt"
)
import "time"

func go_worker(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println("我是一个go协程, 我的名字是 ", name, "----")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, " 执行完毕!")
}

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) //2015-11-16
	//os.Exit(200)
	go go_worker("小黑") //创建一个goroutine协程去执行 go_worker("小黑")
	go go_worker("小白") //创建一个goroutine协程去执行 go_worker("小白")
	go go_worker("小红") //创建一个goroutine协程去执行 go_worker("小白")

	//防止main函数执行完毕,程序退出
	for {
		time.Sleep(1 * time.Second)
	}
}
