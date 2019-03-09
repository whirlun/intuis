package models

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"hdchina/modeldef"
)

func init() {
	orm.RegisterModel(new(modeldef.User))
}

func ValidateUser(name string, passdigest string, hmackey string) error {
	o := orm.NewOrm()
	user := modeldef.User{Name: name}
	err := o.Read(&user, "Name")

	if err == orm.ErrNoRows {
		return errors.New("invalidusername")
	} else if err != nil {
		return errors.New("databaseerror")
	} else {
		h := sha256.New()
		h.Write([]byte("12ws34rf"))
		hmac := hmac.New(sha256.New, []byte(hmackey))
		hmac.Write([]byte(user.PasswordDigest))
		fmt.Println(hmackey)
		fmt.Println(base64.StdEncoding.EncodeToString(h.Sum(nil)))
		fmt.Println(base64.StdEncoding.EncodeToString(hmac.Sum(nil)))
		if p := hmac.Sum(nil); base64.StdEncoding.EncodeToString(p) == passdigest {
			return nil
		} else {
			return errors.New("wrongpassword")
		}
	}
}

func RegisterUser(name string, password string, email string, code string) error {
	o := orm.NewOrm()
	invitation := modeldef.InvitationCode{Code: code}
	err := o.Read(&invitation, "code")
	if err == orm.ErrNoRows {
		return errors.New("invalidinvitation")
	} else if invitation.Used {
		return errors.New("usedcode")
	} else if err != nil {
		return errors.New("databaseerror")
	} else {
		h := sha256.New()
		h.Write([]byte(password))
		passdigest := h.Sum(nil)
		if err = o.Read(&modeldef.User{Name: name}); err != orm.ErrNoRows {
			return errors.New("usedusername")
		}
		user := modeldef.User{Name: name, PasswordDigest: base64.StdEncoding.EncodeToString(passdigest), Email: email}
		mss := modeldef.UserMainSiteSettings{Freeze: false, UseHttps: true,
			Country: "China", Language: "zh_CN",
			FontSize: 20, ShowAdvertisment: true, User: &user}
		fs := modeldef.UserForumSettings{ThreadPerPage: 20, TopicPerPage: 20,
			ShowAvatar: true, ShowSignature: true,
			Signature: "", User: &user}
		ps := modeldef.UserPmSettings{Notify: false, NewTorrent: false,
			NewComment: true, MaxMessage: 20,
			RecieveType: 111, DeleteAfterReply: true,
			SaveSendedMessage: false, ReferedNotice: true, User: &user}
		user.ForumSettings = &fs
		user.MainSiteSettings = &mss
		user.PmSettings = &ps
		err = o.Begin()
		o.Insert(&fs)
		o.Insert(&mss)
		o.Insert(&ps)
		o.Insert(&user)
		invitation.Used = true
		invitation.GuestUser = &user
		o.Update(&invitation, "Used", "GuestUser")
		if err != nil {
			o.Rollback()
			return errors.New("databaseerror")
		} else {
			o.Commit()
			return nil
		}

	}
}

func GetUserPageSetting(username string) (int, error) {
	o := orm.NewOrm()
	u := modeldef.User{Name:username}
	err := o.Read(&u,"Name")
	if err == orm.ErrNoRows {
		return -1, errors.New("invalidusername")
	} else if err != nil {
		return -1, errors.New("databaseerror")
	} else {
		return u.MainSiteSettings.SeedPerPage,nil
	}
}

func ChangeUserForumSetting(username string, settings modeldef.UserForumSettings) error {
	o := orm.NewOrm()
	u := modeldef.User{Name: username}
	err := o.Read(&u, "Name")
	if err == orm.ErrNoRows {
		return errors.New("invalidusername")
	} else if err != nil {
		return errors.New("databaseerror")
	}
	u.ForumSettings = &settings
	oerr := o.Begin()
	o.Update(&u)
	if oerr != nil {
		o.Rollback()
		return errors.New("databaseerror")
	} else {
		o.Commit()
		return nil
	}
}

func ChangeMainSiteSetting(username string, settings modeldef.UserMainSiteSettings) error {
	o := orm.NewOrm()
	u := modeldef.User{Name: username}
	err := o.Read(&u, "Name")
	if err == orm.ErrNoRows {
		return errors.New("invalidusername")
	} else if err != nil {
		return errors.New("databaseerror")
	}
	u.MainSiteSettings = &settings
	oerr := o.Begin()
	o.Update(&u)
	if oerr != nil {
		o.Rollback()
		return errors.New("databaseerror")
	} else {
		o.Commit()
		return nil
	}
}

func ChangePMSettings(username string, settings modeldef.UserPmSettings) error {
	o := orm.NewOrm()
	u := modeldef.User{Name: username}
	err := o.Read(&u, "Name")
	if err == orm.ErrNoRows {
		return errors.New("invalidusername")
	} else if err != nil {
		return errors.New("databaseerror")
	}
	u.PmSettings = &settings
	oerr := o.Begin()
	o.Update(&u)
	if oerr != nil {
		o.Rollback()
		return errors.New("databaseerror")
	} else {
		o.Commit()
		return nil
	}
}