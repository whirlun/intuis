package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"
)

type BaseController struct {
	beego.Controller
	i18n.Locale
	isLogin bool
}

func (this *BaseController) Prepare() {
	if this.setLangVer() {
		this.Redirect(this.Ctx.Request.URL.Path, 302)
	}
	this.Data["mainsiteurl"] = beego.AppConfig.String("site::url")
	this.Data["notseed"] = true
	username := this.GetSession("username")
	if username != nil {
		this.isLogin = true
		this.Data["isLogin"] = true
	} else {
		this.isLogin = false
		this.Data["isLogin"] = false
	}
	var MessageTmpls = map[string]string{
		"Required":     this.Tr("validrequired"),
		"Min":          this.Tr("validmin%d"),
		"Max":          this.Tr("validmax%d"),
		"Range":        this.Tr("validrange"),
		"MinSize":      this.Tr("validminsize"),
		"Length":       this.Tr("validlength"),
		"Alpha":        this.Tr("validalpha"),
		"Numeric":      this.Tr("validnumetic"),
		"AlphaNumeric": this.Tr("validalphanumeric"),
		"Match":        this.Tr("validmatch"),
		"NoMatch":      this.Tr("validnomatch"),
		"AlphaDash":    this.Tr("validalphadash"),
		"Email":        this.Tr("validemail"),
		"IP":           this.Tr("validip"),
		"Base64":       this.Tr("validbase64"),
		"Mobile":       this.Tr("validmobile"),
		"Tel":          this.Tr("validtel"),
		"Phone":        this.Tr("validphone"),
		"ZipCode":      this.Tr("validzipcode"),
	}
	validation.SetDefaultMessage(MessageTmpls)
}

func (this *BaseController) setLangVer() bool {
	isNeedRedir := false
	hasCookie := false

	// 1. Check URL arguments.
	lang := this.Input().Get("lang")

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = this.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := this.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5]
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	if len(lang) == 0 {
		lang = "zh-CN"
		isNeedRedir = false
	}

	if !hasCookie {
		this.Ctx.SetCookie("lang", lang, 1<<31-1, "/")
	}

	this.Lang = lang
	this.Data["Lang"] = lang

	return isNeedRedir
}
