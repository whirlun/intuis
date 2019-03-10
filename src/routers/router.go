package routers

import (
	"github.com/astaxie/beego"
	"controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.SessionController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/newseed", &controllers.SeedController{}, "get:New;post:Post")
	beego.Router("/newseed/image", &controllers.SeedController{}, "post:PostImage")
	beego.Router("/seed", &controllers.SeedController{}, "get:Get")
}
