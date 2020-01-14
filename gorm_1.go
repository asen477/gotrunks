/**
 * User: Trunks(Gao)
 * Date: 2020/01/14
 * gorm框架ORM MySql连接
 */
package main

import (
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//商品
type Food struct {
	Id         int
	Title      string
	Price      float32
	Stock      int
	Type       int
	//mysql datetime, date类型字段，可以和golang time.Time类型绑定， 详细说明请参考：gorm连接数据库章节。
	CreateTime time.Time
}

//为Food绑定表名
func (v Food) TableName() string {
	return "foods"
}

func main(){

	db, err := gorm.Open("mysql", "root:root@/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error()) //错误处理
	}

	//定义接收查询结果的结构体变量
	food := Food{}

	//等价于：SELECT * FROM `foods`   LIMIT 1
	//db.Take(&food)
	db.First(&food,1)

	fmt.Println(food)
}

