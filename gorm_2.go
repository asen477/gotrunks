/**
 * User: Trunks(Gao)
 * Date: 2020/01/15
 * gorm框架ORM MySql连接
 */
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//商品
type Food struct {
	Id    int
	Title string
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

	//定义接收查询结果的结构体变量
	//food := Food{}
	var rows []Food
	//等价于：SELECT * FROM `foods`   LIMIT 1
	//db.Take(&food)
	//db.Find(&food)

	db.Where([]int64{1, 2, 3}).Find(&rows)

	fmt.Println(rows)
	fmt.Println('\n')
	//for i:=0 ;i<10;i++{
	//	db.Find(&rows)
	//	fmt.Println(rows)
	//	time.Sleep(time.Second/10)
	//}
}
