package models

import (
	"github.com/astaxie/beego/orm"
	"hdchina/modeldef"
)


func init() {
	orm.RegisterModel(new(modeldef.IMDBData))
}