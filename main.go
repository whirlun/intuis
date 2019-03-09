package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"hdchina/models"
	_ "hdchina/routers"
	"hdchina/rediscache"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = beego.AppConfig.String("redis::addr")
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.InsertFilter("/*", beego.BeforeRouter, CheckLoginStatus)
	data :=models.InitRedisData()

	rediscache.InitRedis(data)
	/*err := orm.RunSyncdb("default", true, true)
	if err != nil {
		fmt.Println(err)
	}*/
	beego.Run()
}
