/**
 * User: Trunks(Gao)
 * Date: 2020/01/14
 * 9*9乘法表
 */
package main

import (
	"fmt"
	"time"
)

/** 循环公式 **/
func gmsCompute(b int){
	var a int
	for a = 1; a <= b; a++ {
		fmt.Printf("%d*%d=%d ", a,b,a*b)
	}
	fmt.Println("\n")
	time.Sleep(time.Second/5)
}

func main(){
	// 开始执行时间
	timeStart := time.Now()
	var a int
	for a = 1; a < 10; a++ {
		gmsCompute(a)
	}
	// 程序结束时间
	timeEnd := time.Since(timeStart)
	fmt.Println("程序执行时间:",timeEnd)
	fmt.Println("\n9*9乘法表执行完毕!")
	//防止main函数执行完毕,程序退出
	for {
		time.Sleep(1 * time.Second)
	}
}