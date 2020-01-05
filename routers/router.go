package routers

import (
	"github.com/astaxie/beego"
	"github.com/riba2534/OnlineJudge/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
