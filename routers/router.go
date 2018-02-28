package routers

import (
	"github.com/sowjumn/foodie-beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
