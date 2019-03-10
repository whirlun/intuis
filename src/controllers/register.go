package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"models"
	"html/template"
	"strings"
)

type RegisterController struct {
	BaseController
}

type user struct {
	Username              string `form:"username";valid:"Required"`
	Email                 string `form:"email";valid:"Email"`
	Password              string `form:"password";valid:"Required;MinSize(8)"`
	Password_confirmation string `form:"password_confirmation";valid:"Required"`
	Password_question     string `form:"password_question";valid:"Required"`
	Password_answer       string `form:"password_answer";valid:"Required"`
	Invitation_code       string `form:"invitation_code";valid:"Required;Length(16)"`
}

func (this *RegisterController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("Get", this.Get)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

func (this *RegisterController) Get() {
	beego.ReadFromRequest(&this.Controller)
	this.TplName = "session/register.tpl"
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
}

func (this *RegisterController) Post() {
	u := user{}
	if err := this.ParseForm(&u); err != nil {
		beego.Error("Error occured when parse register form: ", err.Error())
		flash := beego.NewFlash()
		flash.Error(this.Tr("internalerror"))
		flash.Store(&this.Controller)
		this.Redirect("/register", 302)
	}
	valid := validation.Validation{}
	b, err := valid.Valid(&u)
	if err != nil {
		beego.Error("validation process error: ", err.Error())
		flash := beego.NewFlash()
		flash.Error(this.Tr("internalerror"))
		flash.Store(&this.Controller)
		this.Redirect("/register", 302)
	}
	if !b {
		flashmessage := ""
		for _, err := range valid.Errors {
			flashmessage = flashmessage + err.Key + err.Message + "<br />"
		}
		flash := beego.NewFlash()
		flash.Error(flashmessage)
		flash.Store(&this.Controller)
		this.Redirect("/register", 302)
	}
	if strings.Compare(u.Password, u.Password_confirmation) != 0 {
		flash := beego.NewFlash()
		flash.Error("passwordinconsistent")
		flash.Store(&this.Controller)
		this.Redirect("/register", 302)
	}
	rerr := models.RegisterUser(u.Username, u.Password, u.Email, u.Invitation_code)
	if rerr != nil {
		flash := beego.NewFlash()
		flash.Error(this.Tr(rerr.Error()))
		flash.Store(&this.Controller)
		this.Redirect("/register", 302)
	} else {
		this.SetSession("username", u.Username)
		this.Redirect("/", 302)
	}
}
