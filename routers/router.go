package routers

import (
	"github.com/mikelut76/bee-demo-web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
