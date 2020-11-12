package routers

import (
	"hello2/controllers"

	beego "github.com/astaxie/beego/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/users", &controllers.UserController{})

	beego.SetStaticPath("/down2", "static2")

}
