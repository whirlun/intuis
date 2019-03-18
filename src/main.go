package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/beego/i18n"
	"models"
	"rediscache"
	_ "routers"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = beego.AppConfig.String("redis::addr")
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.InsertFilter("/*", beego.BeforeRouter, CheckLoginStatus)
	data := models.InitRedisData()

	rediscache.InitRedis(data)
	/*err := orm.RunSyncdb("default", true, true)
	if err != nil {
		fmt.Println(err)
	}*/
	beego.Run()
}
