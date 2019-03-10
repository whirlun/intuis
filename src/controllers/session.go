package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"models"
	"html/template"
	"math/rand"
	"strconv"
	"time"
)

type SessionController struct {
	BaseController
}

type login struct {
	Username string `valid:"Required"`
	Password string `valid:"Required"`
}

func (this *SessionController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("Get", this.Get)
}

func (this *SessionController) Post() {
	u := login{Username: this.GetString("username"), Password: this.GetString("password")}
	valid := validation.Validation{}
	b, err := valid.Valid(&u)
	if err != nil {
		beego.Error("validation process error: ", err.Error())
		flash := beego.NewFlash()
		flash.Error(this.Tr("internalerror"))
		flash.Store(&this.Controller)
		this.Redirect("/register", 200)
	}
	if !b {
		jsondata := struct {
			Reason string
		}{this.Tr("logininvalid")}
		this.Data["json"] = jsondata
		this.ServeJSON()
	}
	hmackey := this.GetString("_hmackey")
	err = models.ValidateUser(u.Username, u.Password, hmackey)
	if err != nil {
		jsondata := struct {
			Reason string
		}{this.Tr(err.Error())}
		this.Data["json"] = jsondata
		this.ServeJSON()
	} else {
		this.SetSession("username", u.Username)
		jsondata := struct {
			Reason string
		}{"success"}
		this.Data["json"] = jsondata
		this.ServeJSON()
	}
}

func (this *SessionController) Get() {
	beego.ReadFromRequest(&this.Controller)
	this.TplName = "session/login.tpl"
	timestamp := time.Now().Unix()
	rand.Seed(time.Now().UnixNano())
	hmackey := (uint(timestamp)<<(uint(rand.Intn(255))>>4) ^ (uint(rand.Intn(255)) << 4))
	this.Data["hmackey"] = strconv.FormatUint(uint64(hmackey), 10)
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.Data["xsrftoken"] = this.XSRFToken()
}
