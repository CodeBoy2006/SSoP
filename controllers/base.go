package controllers

import (
	"github.com/astaxie/beego"

	"fmt"
)

type BaseController struct {
	beego.Controller
	Uid string
	Gid string
}

func (c *BaseController) Prepare() {
	sess := c.GetSession("uid")
	beego.Debug(c.IsAjax())
	if sess == nil {
		if !c.IsAjax() {
			c.Ctx.Redirect(302, "/store/login")
			return
		} else {
			c.Data["json"] = "wdl"
			c.ServeJSON()
		}
	}

	c.Uid = fmt.Sprint(sess)
}
