package main

import (
	"github.com/astaxie/beego/context"
)

func CheckLoginStatus(ctx *context.Context) {
	if ctx != nil {
		if !(ctx.Request.URL.Path == "/register" || ctx.Request.URL.Path == "/login") {
			_, ok := ctx.Input.Session("username").(string)
			if !ok && ctx.Request.RequestURI != "/login" {
				ctx.Redirect(302, "/login")
			}
		}
	}
}
