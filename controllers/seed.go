package controllers

import (
	"bytes"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/satori/go.uuid"
	"gopkg.in/h2non/filetype.v1"
	"hdchina/models"
	"hdchina/rediscache"
	"io"
	"regexp"
	"time"
	"os"
	"strings"
	"fmt"
	"html/template"
	"crypto/sha1"

	"encoding/hex"
	"github.com/astaxie/beego/utils/pagination"
	"hdchina/modeldef"
)

const (
	MAX_IMAGE_SIZE = 1 << 22
	IMAGE_TYPES     = "(jpg|webp|jpeg|png)"
)

var (
	imageTypes = regexp.MustCompile(IMAGE_TYPES)
	nowtime = time.Now().Format("2006-01-02_150405")

)

type SeedController struct {
	BaseController
}

type seed struct {
	Title           string `form:"title";valid:"Required"`
	Subtitle        string `form:"substring";valid:"Required`
	Is3D            bool   `form:"is3D";valid:"Required"`
	ReferRule       string `form:"referrule";valid:"Required"`
	Medium          string `form:"medium";valid:"Required"`
	VideoEncode     string `form:"videoencode";valid:"Required"`
	AudioEncode     string `form:"audioencode";valid:"Required"`
	Category        string `form:"category";valid:"Required"`
	ProductionGroup string `form:"productiongroup";valid:"Required"`
	Format          string `form:"format";valid:"Required"`
	IMDBLink        string `form:"imdblink";valid:"Required"`
	DoubanLink      string `form:"doubanlink";valid:"Required"`
	Content         string `form:"content-markdown-doc";valid:"Required"`
}

type imagejsondata struct {
	Success int `json:"success"`
	Message string `json:"message"`
	Url string `json:"url"`
}

func (this *SeedController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("New", this.New)
}

//在简介中上传图片的处理函数
func (this *SeedController) PostImage() {
	imf, ih, _ := this.GetFile("editormd-image-file") //获取图片文件
	username := this.GetSession("username").(string) //获取用户名，图片将存在用户名为名称的目录下
	if exist, _ := pathExists("static/upload/content/" + username); !exist { //检查目录是否存在
		err := os.MkdirAll("static/upload/content/"+username, 0777) //不存在则创建
		if err != nil { //如果创建失败则返回系统内部错误信息
			beego.Error("error in creating user folder", err.Error())
			jsondata := imagejsondata{Success: 0, Message: this.Tr("internalerror")}
			this.Data["json"] = jsondata
			this.ServeJSON()
		}
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.CopyN(buf, imf,261); err != nil { //获取图片文件头
		beego.Error("error occured when checking file type ", err)
		jsondata := imagejsondata{Success: 0, Message: this.Tr("internalerror")}
		this.Data["json"] = jsondata
		this.ServeJSON()
	}
	head := buf.Bytes()
	ptype, _ := filetype.Match(head) //检查文件类型是否允许
	if !imageTypes.MatchString(ptype.Extension) {
		jsondata := imagejsondata{Success: 0, Message: this.Tr("postertypenotaccepted")}
		this.Data["json"] = jsondata
		this.ServeJSON()
	}
	if ih.Size > MAX_IMAGE_SIZE { //检查图片大小是否超出允许范围
		jsondata := imagejsondata{Success: 0, Message: this.Tr("postersizenotaccepted")}
		this.Data["json"] = jsondata
		this.ServeJSON()
	}
	var ctfilename []string
	if this.GetSession("contentimages") != nil { //从session中获取已经上传的图片列表
		ctfilename = this.GetSession("contentimages").([]string)
	}
	h := sha1.New()
	_, err := io.Copy(h, imf) //使用SHA1创建图片文件名
	if err != nil {
		beego.Error("error in creating sha1 value")
		jsondata := imagejsondata{Success: 0, Message: this.Tr("internalerror")}
		this.Data["json"] = jsondata
		this.ServeJSON()
	}
	content_path := "static/upload/content/" + username + "/" + hex.EncodeToString(h.Sum(nil))
	var url string
	if exist, _ := pathExists(content_path); !exist { //如果文件重复则直接返回已有文件
		err := this.SaveToFile("editormd-image-file", content_path)
		fmt.Println(err)
		ctfilename = append(ctfilename, content_path)
		}

	this.SetSession("contentimages", ctfilename) //将新文件加入session
	url = "http://"+beego.AppConfig.String("site::url")+"/static/upload/content/" + username + "/" + hex.EncodeToString(h.Sum(nil))
	jsondata := imagejsondata{1, this.Tr("success"),url}
	this.Data["json"] = jsondata
	this.ServeJSON()
	}
func (this *SeedController) Post() {
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.Data["csrftoken"] = this.XSRFToken()
	this.Data["notseed"] = false
	this.TplName = "seed/new.tpl"
	this.Data["editor"] = true
	s := seed{}
	if err := this.ParseForm(&s); err != nil { //获取form信息
		this.internalerror("error occured when parsing form", err)
	}
	valid := validation.Validation{}
	b, err := valid.Valid(&s) //检查form
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
	username := this.GetSession("username").(string) //创建文件保存目录
	posterpath := "static/upload/poster/" +username + "_" + uuid.Must(uuid.NewV1()).String() + "_" + nowtime
	torrentpath := "static/upload/torrent/"+ username+ "_" + uuid.Must(uuid.NewV1()).String() + "_" + nowtime + ".torrent"
	nfopath := "static/upload/nfo/" + username  + uuid.Must(uuid.NewV1()).String() + "_" + nowtime + ".nfo"
	tf, th, _:= this.GetFile("torrent") //如果没有上传种子就返回错误信息
	if th.Size == 0 {
		this.missingfile(this.Tr("torrent"))
	} else {
		this.SaveToFile("torrent", torrentpath)
	}
	pf, ph, _ := this.GetFile("poster") //如果上传了海报就进行保存
	if th.Size == 0 {
		posterpath = ""
	} else {
		buf := bytes.NewBuffer(nil)
		if _, err := io.CopyN(buf, pf,261); err != nil { //获取文件头
			this.internalerror("error occured when checking file type ", err)
		}
		head := buf.Bytes()
		ptype, _ := filetype.Match(head)
		if !imageTypes.MatchString(ptype.Extension) { //检查是否是允许的图片类型，不允许则返回错误
			flash := beego.NewFlash()
			flash.Error(this.Tr("postertypenotaccepted"))
			flash.Store(&this.Controller)
			this.Redirect("/newseed", 302)
		}
		if ph.Size > MAX_IMAGE_SIZE { //检查文件大小是否合适
			flash := beego.NewFlash()
			flash.Error(this.Tr("postersizenotaccepted"))
			flash.Store(&this.Controller)
			this.Redirect("/newseed", 302)
		}
		this.SaveToFile("poster", posterpath)
	}
	nf, nh, _ := this.GetFile("nfo") //如果上传了nfo文件就保存
	if nh.Size == 0 {
		nfopath = ""
	} else {
		this.SaveToFile("nfo", nfopath)
	}
	defer tf.Close() //关闭打开的文件句柄
	defer pf.Close()
	defer nf.Close()
	var ctfilename []string
	if this.GetSession("contentimages") != nil { //获取所有在简介中上传的图片
		ctfilename = this.GetSession("contentimages").([]string)
	}
	ss := modeldef.Seed{
		Title:      s.Title,
		Subtitle:   s.Subtitle,
		DoubanLink: s.DoubanLink,
		IMDBLink:   s.IMDBLink,
		Poster:     posterpath,
		Torrent:    torrentpath,
		Nfo:        nfopath,
		Is3D:       s.Is3D,
		Top:        false,
		Locked:     false,
		Content:    s.Content,
		ContentImage: strings.Replace(strings.Trim(fmt.Sprint(ctfilename), "[]"), " ", ",", -1),
	}
	serr := models.SaveSeed(ss, this.GetSession("username").(string), s.Category, s.Format, s.ProductionGroup, s.Medium, s.VideoEncode, s.AudioEncode, s.ReferRule)
	if serr != nil {
		flash := beego.NewFlash()
		flash.Error(this.Tr(serr.Error()))
		flash.Store(&this.Controller)
		this.Redirect("/newseed", 302)
	}
}

func (this *SeedController) New() {
	if this.GetSession("contentimages") != nil && len(this.GetSession("contentimages").([]string)) != 0 {
		this.DelSession("contentimages")
	}
	beego.ReadFromRequest(&this.Controller)
	ch := make(chan map[string][]string)
	go rediscache.GetSeedProperties(ch)
	var propmap map[string] []string
	propmap = <-ch
	if len(propmap) == 0 {
		go models.GetPropertyData(ch)
		propmap = <-ch
	}
	for dbname, properties := range propmap {
		this.Data[dbname] = properties
	}
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.Data["csrftoken"] = this.XSRFToken()
	this.Data["notseed"] = false
	this.TplName = "seed/new.tpl"
	this.Data["editor"] = true
}

func (this *SeedController) Get() {
	this.TplName = "seed/index.tpl"
	cnt, _ := models.GetSeedCount()
	useed, _ := models.GetUserPageSetting(this.GetSession("username").(string))
	paginator := pagination.SetPaginator(this.Ctx,useed,cnt)
	this.Data["paginator"] = paginator
	_type := this.GetString("type")
	startpage,err  := this.GetInt("start")
	if err != nil {
		this.Redirect("/seed", 302)
	}
	seedlist, err := models.GetSeedList(startpage, paginator.PerPageNums, _type)
	if err != nil {
		this.internalerror("error occured when reading seedlist",err)
	}
	seeds ,err := models.GetSeedPropertyById(seedlist)
	if err != nil {
		beego.Error("error occured when getting seeds from database/redis",err.Error())
		flash := beego.NewFlash()
		flash.Error(this.Tr(err.Error()))
		flash.Store(&this.Controller)
		this.Redirect("/seed", 302)
	}
	this.Data["seeds"] = seeds
}

func (this *SeedController) internalerror(msg string, err error) {
	beego.Error(msg, err.Error())
	flash := beego.NewFlash()
	flash.Error(this.Tr("internalerror"),err.Error())
	flash.Store(&this.Controller)
	this.Redirect("/newseed", 302)
}

func (this *SeedController) missingfile(msg string) {
	beego.Error(msg, "missing upload file:" + msg)
	flash := beego.NewFlash()
	flash.Error(this.Tr("missingfile",msg))
	flash.Store(&this.Controller)
	this.Redirect("/newseed", 302)
}

func pathExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return true, nil
	}
}