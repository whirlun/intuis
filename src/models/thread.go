package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"modeldef"
	"time"
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
	user := modeldef.User{Name: username}
	err = o.Read(&user, "name")
	if err == orm.ErrNoRows {
		return errors.New("invaliduser")
	} else if err != nil {
		return errors.New("databaseerror")
	}
	currenttime := time.Now().Unix()
	thread := modeldef.Thread{Title: title,
		Content:       content,
		PostTime:      currenttime,
		LastReplyTime: 0,
		Reply:         0,
		Read:          0,
		Pinned:        false,
		Locked:        false,
		Category:      &cat,
		User:          &user,
		Replies:       nil,
		LastReplyUser: nil,
	}
	oerr := o.Begin()
	o.Insert(&thread)
	if oerr != nil {
		o.Rollback()
		beego.Error("error occured when commiting database in saving seed", err.Error())
		return errors.New("databaseerror")
	} else {
		o.Commit()
		return nil
	}
}

func LoveThread(tid int, uid int, tuid int) error {
	//alternative way to judge whether a row exists in order to avoid future performance problem:
	//select isnull((select top(1) 1 from tableName where conditions), 0)
	o := orm.NewOrm()
	thread := modeldef.Thread{Id: tid}
	err := o.Read(&thread, "id")
	if err == orm.ErrNoRows {
		return errors.New("invalidthread")
	} else if err != nil {
		return errors.New("databaserror")
	}
	user := modeldef.User{Id: uid}
	err = o.Read(&user, "id")
	if err == orm.ErrNoRows {
		return errors.New("invaliduser")
	} else if err != nil {
		return errors.New("databaseerror")
	}
	tuser := modeldef.User{Id: tuid}
	err = o.Read(&tuser, "id")
	if err == orm.ErrNoRows {
		return errors.New("invaliduser")
	} else if err != nil {
		return errors.New("databaseerror")
	}
	loves := modeldef.Loves{Operator: &user, Receiver: &tuser, Thread: &thread}
	lc := modeldef.LovesCount{Operator: user.Name, Receiver: tuser.Name}
	oerr := o.Begin()
	o.Raw("UPDATE thread SET love=love+1 WHERE id =" + string(tid)).Exec()
	o.Insert(&loves)
	created, id, _ := o.ReadOrCreate(&lc, "operator", "receiver")
	if created {
		o.Update(&modeldef.LovesCount{Id: int(id), Count: 1}, "count")
	} else {
		o.Update(&modeldef.LovesCount{Id: int(id), Count: lc.Count + 1}, "count")
	}
	if oerr != nil {
		o.Rollback()
		beego.Error("error occured when commiting database in saving love of the thread", err.Error())
		return errors.New("databaseerror")
	} else {
		o.Commit()
	}
	return nil
}

func DeloveThread(tid int, uid int, tuid int) error {
	o := orm.NewOrm()
	thread := modeldef.Thread{Id: tid}
	err := o.Read(&thread, "id")
	if err == orm.ErrNoRows {
		return errors.New("invalidthread")
	} else if err != nil {
		return errors.New("databaserror")
	}
	user := modeldef.User{Id: uid}
	err = o.Read(&user, "id")
	if err == orm.ErrNoRows {
		return errors.New("invaliduser")
	} else if err != nil {
		return errors.New("databaseerror")
	}
	tuser := modeldef.User{Id: tuid}
	err = o.Read(&tuser, "id")
	if err == orm.ErrNoRows {
		return errors.New("invaliduser")
	} else if err != nil {
		return errors.New("databaseerror")
	}
	loves := modeldef.Loves{Operator: &user, Receiver: &tuser, Thread: &thread}
	lc := modeldef.LovesCount{Operator: user.Name, Receiver: tuser.Name}
	oerr := o.Begin()
	o.Raw("UPDATE thread SET love=love-1 WHERE id =" + string(tid)).Exec()
	o.Delete(&loves)
	created, id, _ := o.ReadOrCreate(&lc, "operator", "receiver")
	if created {
	} else {
		o.Update(&modeldef.LovesCount{Id: int(id), Count: lc.Count - 1}, "count")
	}
	if oerr != nil {
		o.Rollback()
		beego.Error("error occured when commiting database in saving love of the thread", err.Error())
		return errors.New("databaseerror")
	} else {
		o.Commit()
	}
	return nil
}
