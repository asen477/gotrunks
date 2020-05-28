/**
 * User: Trunks(Gao)
 * Date: 2020/01/16
 * 活动抽奖规则
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 初始化库存
var arr_number_data = [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
var arr_zodiac_data = [12]string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}

// 随机数结构体
type randItem struct {
	start float64
	end   float64
}

func main() {
	//arr_number_data[0] = arr_number_data[0] - 1
	//arr_number_data[11] = arr_number_data[11] - 9
	//setStockDeduct(11)

	fmt.Println("当前生肖卡库存：", arr_number_data)
	fmt.Println(arr_zodiac_data)

	//os.Exit(200)

	// 执行抽卡
	res := GetSSRcard(10)
	fmt.Println("我的到的SSR卡片：", arr_zodiac_data[res])

	// 0~15（出现的概率为50%），16~31（出现的概率为20%），32~64（出现的概率为20%），64~128（出现的概率为10%）
	probabilityPrint()
}

// GetSSRcard 获得随机数算法
func GetSSRcard(num int) int {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < num; i++ {
		fmt.Println("随机生成：", rand.Intn(12))
		//fmt.Println("\n")
	}
	v := rand.Intn(12)
	fmt.Println("返回抽卡结果：", v)
	return v
}

// GetRandItem TODO GetRandItem 概率抽奖算法
func GetRandItem(items map[interface{}]float64) interface{} {

	nums := make(map[interface{}]randItem)

	var s float64

	for k, f := range items {
		if f <= 0 {
			continue
		}
		nums[k] = randItem{
			start: s,
			end:   s + f,
		}
		s = nums[k].end
	}

	rnd := rand.Float64() * s

	for k, f := range nums {
		if f.start <= rnd && f.end > rnd {
			return k
		}
	}

	return nil
}

// probabilityPrint 概率调用输出
func probabilityPrint() {

	rand.Seed(time.Now().UnixNano())

	type info struct {
		Start int
		End   int
	}

	mItems := map[interface{}]float64{
		info{
			Start: 0,
			End:   15,
		}: 0.5,
		info{
			Start: 16,
			End:   31,
		}: 0.2,
		info{
			Start: 32,
			End:   64,
		}: 0.2,
		info{
			Start: 65,
			End:   128,
		}: 0.1,
	}

	item := GetRandItem(mItems).(info)
	fmt.Println("===", item)
}

// 库存扣减
func setStockDeduct(id int) int {
	arr_number_data[id] = arr_number_data[id] - 1
	return arr_number_data[id]
}

// 库存增加
func setStockAdd(id int) int {
	arr_number_data[id] = arr_number_data[id] + 1
	return arr_number_data[id]
}
