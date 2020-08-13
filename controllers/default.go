package controllers

import (
	"day3/models"
	//"encoding/json"
	"github.com/astaxie/beego"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Store() {
	plist := models.Query("select * from products")
	c.Data["product"] = plist
	beego.Debug(plist)
	c.TplName = "homepage.html"

}
func (c *MainController) Products() {
	plist := models.Query("select * from products")
	c.Data["product"] = plist
	beego.Debug(plist)
	c.TplName = "homepage.html"
}

type Register struct {
	beego.Controller
}

func (c *Register) Post() {
	username := c.GetString("username")
	password := models.ToMD5(c.GetString("password"))
	phone := c.GetString("phone")
	email := c.GetString("email")
	_, num := models.Querynum(`select id from users where email=?`, email)
	if num == 0 {
		err := models.Exec(`insert into users (username,password,phone,email)
							values (?,?,?,?)`, username, password, phone, email)
		if err != nil {
			beego.Debug(err)
			c.Data["json"] = 0
		} else {
			c.Data["json"] = 1
		}
	} else {
		c.Data["json"] = 3
	}

	c.ServeJSON()
}
func (c *Register) Get() {
	c.TplName = "register.html"
}

type Login struct {
	beego.Controller
}

func (c *Login) Post() {
	password := models.ToMD5(c.GetString("password"))
	email := c.GetString("email")

	p := models.Query(`select * from users where email=? limit 1`, email)

	if p == nil {
		c.Data["json"] = 0
	} else {

		if password == p[0]["password"] {
			c.Data["json"] = 1
			c.SetSession("uid", p[0]["id"])
		} else {
			c.Data["json"] = 2
		}
	}
	c.ServeJSON()
}
func (c *Login) Get() {
	c.TplName = "login.html"
}

func (c *MainController) Category() {
	c.TplName = "category.html"
}

func (c *MainController) Cart() {
	c.TplName = "cart.html"
}

func (c *MainController) My() {
	c.TplName = "my.html"
}

func (c *MainController) Detail() {
	c.TplName = "detail.html"
}

func (c *MainController) BuyP() {
	c.TplName = "buy.html"
}

func (c *MainController) Changeavatar() {
	c.TplName = "changeAvatar.html"
}
