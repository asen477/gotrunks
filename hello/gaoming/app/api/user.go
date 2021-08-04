package api

import (
	"github.com/gogf/gf/net/ghttp"
)

var User = userApi{}

type userApi struct {}

func (*userApi) Index(r *ghttp.Request){
	r.Response.Writeln("abc!")
}