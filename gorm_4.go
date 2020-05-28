/**
 * User: Trunks(Gao)
 * Date: 2020/05/25
 * gorm框架ORM MySql创建记录
 * URL链接：https://gorm.io/
 */

package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Foods 记录
type Foods struct {
	Id    uint   `gorm:"primary_key"`
	Title string `gorm:"default:'galeone'"`
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error()) //错误处理
	}

	var rows []Foods
	db.First(&rows)

	fmt.Println(rows)
}
