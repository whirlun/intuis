package controllers

type IndexController struct {
	BaseController
}

func (this *IndexController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("Get", this.Get)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

func (this *IndexController) Post() {

}

func (this *IndexController) Get() {
	this.TplName = "index/index.tpl"
	this.Data["image1"] = "/static/img/1.png"
	this.Data["image2"] = "/static/img/2.png"
	this.Data["image3"] = "/static/img/a3.png"
	this.Data["image4"] = "/static/img/a4.png"
	this.Data["image5"] = "/static/img/a5.png"
}

func (this *IndexController) Put() {

}

func (this *IndexController) Delete() {

}
