/**
 * User: Trunks(Gao)
 * Date: 2020/05/28
 * 指针(pointer)
 */
package main

import "fmt"

func add(num int) {
	num += 1
	fmt.Println(num)
}

func realAdd(num *int) {
	*num += 1
}

func main() {
	num := 100
	add(num)
	fmt.Println(num) // 100，num 没有变化

	realAdd(&num)
	fmt.Println(num) // 101，指针传递，num 被修改

	// ==========================================================
	str := "Golang"
	var p *string = &str // p 是指向 str 的指针
	*p = "Hello"
	fmt.Println(str) // Hello 修改了 p，str 的值也发生了改变
}
