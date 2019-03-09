package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"fmt"
)

type SeedProperty interface {
	GetString() string
}

var (
	DatabaseNames = []string{"category", "format", "production_group",
		"medium", "video_encode", "audio_encode", "refer_rule"}
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:12ws34rf@/hdchina?charset=utf8mb4")
}

func GetPropertyData(ch chan map[string] []string) {
	o := orm.NewOrm()
	result := make(map[string][]string)
	for _, dbname := range DatabaseNames {
		seedproperty := make([]string, 0, 1)
		if num, err := o.Raw("SELECT " + dbname + " FROM " + dbname).QueryRows(&seedproperty); err == nil {
			sseedproperty := make([]string, num, num)
			for _, property := range seedproperty {
				sseedproperty = append(sseedproperty, property)
			}
			result[dbname] = sseedproperty
		} else {
			fmt.Println(err)
		}
	}
	ch <- result
}

func InitRedisData() map[string] []string{
	ch := make(chan map[string] []string)
	go GetPropertyData(ch)
	return <- ch
}
