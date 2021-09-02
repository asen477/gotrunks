package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func main(){
	r,err := g.Model("user").All()
	fmt.Print(err)
	fmt.Print(r)
	s := g.Server()
	// 测试日志
	s.BindHandler("/welcome", func(r *ghttp.Request) {
		glog.Info("你来了！")
		glog.Error("你异常啦！")
		r.Response.Write("哈喽世界！")
	})
	// 异常处理
	s.BindHandler("/panic", func(r *ghttp.Request) {
		glog.Panic("123")
	})
	// post请求
	s.BindHandler("POST:/hello", func(r *ghttp.Request) {
		r.Response.Writeln("Hello World!")
	})
	// BindHandler是最原生的路由注册方法，在大部分场景中，我们通常使用 分组路由 方式来管理路由
	s.BindHandler("GET:/hello", func(r *ghttp.Request) {
		r.Response.Writeln("url" + r.Router.Uri)
	})
	// 分组注册及中间件
	group := s.Group("/hello")
	group.ALL("/all", func(r *ghttp.Request) {
		r.Response.Writeln("all")
	})

	s.BindHandler("GET:/gao", func(r *ghttp.Request) {
		s := g.Server()
		s.SetIndexFolder(true)
		s.SetServerRoot("/")
	})

	s.Run()
}