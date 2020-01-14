/**
 * User: Trunks(Gao)
 * Date: 2019/11/09
 * 函数判断是否是质数
 */
package main

import (
	"errors"
	"fmt"
	"math"
)

func main(){
	fmt.Print(IfPrime(77))
}

//函数判断是否是质数
func IfPrime(value int) (bool, error) {
	if value < 2 {
		return false, errors.New("小于2的整数无法判断")
	}
	end := int(math.Sqrt(float64(value)))
	for i := 2; i <= end; i++ {
		if value%i == 0 {
			return false, errors.New("不能被自己整除")
		}
	}
	return true, errors.New("是质数")
}
