package controllers

import (
	"html/template"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type ForumController struct {
	BaseController
}

type thread struct {
	Title string `form:"title";valid:"Required"`
	Content string `form:"title";valid:"Required"`
	Category string `form:"category";valid:"Required"`
}

func (this *ForumController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("Get", this.Get)
}

func (this *ForumController) Post() {
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.Data["csrftoken"] = this.XSRFToken()
	this.TplName = "forum/new.tpl"
	this.Data["editor"] = true
	t := thread{}
	if err := this.ParseForm(&t); err != nil {
		this.internalerror("error occured when parsing form", err)
	}
	valid := validation.Validation{}
	b, err := valid.Valid(&t)
	if err != nil {
		this.internalerror("validation process error: ", err)
	}
	if !b { //如果有错误就返回错误信息
		flashmessage := ""
		for _, err := range valid.Errors {
			flashmessage = flashmessage + err.Key + err.Message + "<br />"
		}
		flash := beego.NewFlash()
		flash.Error(flashmessage)
		flash.Store(&this.Controller)
		this.Redirect("/newseed", 302)
	}

}

func (this *ForumController) Get() {

}

func (this *ForumController) internalerror(msg string, err error) {
	beego.Error(msg, err.Error())
	flash := beego.NewFlash()
	flash.Error(this.Tr("internalerror"),err.Error())
	flash.Store(&this.Controller)
	this.Redirect("/newthread", 302)
}

func (this *ForumController) missingfile(msg string) {
	beego.Error(msg, "missing upload file:" + msg)
	flash := beego.NewFlash()
	flash.Error(this.Tr("missingfile",msg))
	flash.Store(&this.Controller)
	this.Redirect("/newthread", 302)
}