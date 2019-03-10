package models

import (
	"github.com/astaxie/beego/orm"
	"modeldef"
	"errors"
	"time"
	"github.com/astaxie/beego"
)

func init() {
	orm.RegisterModel(new(modeldef.Thread))
}

func PostThread(title string, content string, category string, username string) error {
	o := orm.NewOrm()
	cat := modeldef.ForumCategory{Name: category}
	err := o.Read(&cat, "name")
	if err == orm.ErrNoRows {
		return errors.New("invalidcat")
	} else if err != nil {
		return errors.New("databaseerror")
	}
	user := modeldef.User{Name:username}
	err = o.Read(&user, "name")
	if err == orm.ErrNoRows{
		return errors.New("invaliduser")
	} else if err != nil {
		return errors.New("databaseerror")
	}
	currenttime := time.Now().Unix()
	thread := modeldef.Thread{Title: title,
	Content: content,
	PostTime:currenttime,
	LastReplyTime:0,
	Reply:0,
	Read:0,
	Pinned: false,
	Locked: false,
	Category:&cat,
	User:&user,
	Replies:nil,
	LastReplyUser:nil,
	}
	oerr := o.Begin()
	o.Insert(&thread)
	if oerr != nil {
		o.Rollback()
		beego.Error("error occured when commiting database in saving seed", err.Error())
		return errors.New("internalerror")
	} else {
		o.Commit()
		return nil
	}
}

