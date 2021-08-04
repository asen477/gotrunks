package router

import (
	"gaoming/app/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/hello", api.Hello)
	})
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/user", api.User)
	})
}
