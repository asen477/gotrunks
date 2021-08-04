package main

import (
	_ "gaoming/boot"
	_ "gaoming/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
