package main

import (
	"database/sql"
	"fmt"
	"github.com/gogf/gf/frame/g"
)

//数据结构
type Userinfo struct {
	Id  int
	CnName string
	Telphones string
}

func main() {
	// 数量
	count, err := g.DB().Model("tb_userinfo").Where("UserType", 0).One()
	if err != nil {
		panic(err)
	}
	fmt.Println("姓名:", count.Map()["CnName"])
	fmt.Println("电话:", count.Map()["Telphones"])

	// Record记录处理
	userinfo := (*Userinfo)(nil)
	err2  := g.DB().Table("tb_userinfo").Where("Id=?", 27).Scan(&userinfo)
	if err2 != nil && err2 != sql.ErrNoRows {
		fmt.Println(err2)
	}
 
	if userinfo != nil {
		fmt.Println(" uid:", userinfo.Id)
		fmt.Println("name:", userinfo.CnName)
		fmt.Println("Telphones:", userinfo.Telphones)
	}

}