package routers

import (
	"controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.SessionController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/newseed", &controllers.SeedController{}, "get:New;post:Post")
	beego.Router("/newseed/image", &controllers.SeedController{}, "post:PostImage")
	beego.Router("/seed", &controllers.SeedController{}, "get:Get")
	beego.Router("/user", &controllers.UserController{})
}
