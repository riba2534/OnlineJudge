package routers

import (
	"github.com/riba2534/OnlineJudge/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
