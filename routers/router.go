package routers

import (
	"github.com/astaxie/beego"
	"github.com/flyray/myblog/controllers"
)

func init() {
	//注解路由 需要调用Include。
	beego.Include(&controllers.IndexController{})
}
