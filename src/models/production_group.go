package models

import (
	"github.com/astaxie/beego/orm"
	"modeldef"
)


func init() {
	orm.RegisterModel(new(modeldef.ProductionGroup))
}

