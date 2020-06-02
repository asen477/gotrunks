/**
 * User: Trunks(Gao)
 * Date: 2020/05/28
 * 切片
 */
package main

import "fmt"

func main() {
	//slice1 := make([]float32, 0) // 长度为0的切片
	slice2 := make([]float32, 3, 5)       // [0 0 0] 长度为3容量为5的切片
	fmt.Println(len(slice2), cap(slice2)) // 3 5

	// 添加元素，切片容量可以根据需要自动扩展
	slice2 = append(slice2, 1, 2, 3, 4)   // [0, 0, 0, 1, 2, 3, 4]
	fmt.Println(len(slice2), cap(slice2)) // 7 12
	// 子切片 [start, end)
	sub1 := slice2[3:]  // [1 2 3 4]
	sub2 := slice2[:3]  // [0 0 0]
	sub3 := slice2[1:4] // [0 0 1]
	// 合并切片
	combined := append(sub1, sub2...)      // [1, 2, 3, 4, 0, 0, 0]
	combined2 := append(combined, sub3...) // [1 2 3 4 0 0 0 0 0 1]
	fmt.Println(combined2)
}
