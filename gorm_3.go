/**
 * User: Trunks(Gao)
 * Date: 2020/05/25
 * gorm框架ORM MySql创建记录
 * URL链接：https://books.studygolang.com/gorm/
 */

package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// Foods 商品
type Foods struct {
	Id    uint   `gorm:"primary_key"`
	Title string `gorm:"default:'galeone'"`
	Price float32
	Stock int
	Type  int
	//mysql datetime, date类型字段，可以和golang time.Time类型绑定， 详细说明请参考：gorm连接数据库章节。
	CreateTime time.Time
}

func main() {

	db, err := gorm.Open("mysql", "root:root@/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error()) //错误处理
	}

	// TODO Golang 获取当前表的记录总数
	var count int
	db.Raw("SELECT count(*) From foods").Count(&count)
	fmt.Println("插入记录前：", count)

	// TODO Golang 创建记录
	var data = &Foods{Title: "L1212", Price: 1000, Stock: 9090, Type: 2, CreateTime: time.Now()}
	db.Create(data)
	fmt.Println(time.Now().Date())

	// TODO 统计记录
	db.Raw("SELECT count(*) From foods").Count(&count)
	fmt.Println("插入记录后：", count)

	// TODO 查询
	var rows []Foods
	db.First(&rows, 1) // 查询id为1的product
	fmt.Println(rows)

	// TODO 更新 - 更新rows的price为2020
	db.Model(&rows).Where("id = ?", 1).Update("Price", 2020)
	fmt.Println(rows)

	defer db.Close()
}
