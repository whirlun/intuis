package routers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"strings"
)

func init() {
	langs := strings.Split(beego.AppConfig.String("lang"), "|")
	for _, lang := range langs {
		lang = strings.TrimSpace(lang)
		if err := i18n.SetMessage(lang, "conf/i18n/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Failed to set i18n file:" + err.Error())
			return
		}
	}
}
