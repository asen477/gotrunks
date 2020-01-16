/**
 * User: Trunks(Gao)
 * Date: 2020/01/16
 * 函数包
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 获得随机数量
func getSSRcard(num int) int{
	rand.Seed(time.Now().UnixNano())
	for i:=1;i<num;i++{
		fmt.Println("随机生成：",rand.Intn(12))
		fmt.Println("\n")
	}
	v := rand.Intn(12)
	return v
}

// SetStockDeduct 库存扣减
func SetStockDeduct(id int) int{
	return id
}

// SetStockAdd 库存增加
func SetStockAdd(id int) int{
	return id
}