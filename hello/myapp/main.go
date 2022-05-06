package main

import (
	_ "myapp/boot"
	_ "myapp/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
