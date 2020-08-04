package controllers

import (
	"day3/models"
	//"encoding/json"
	"github.com/astaxie/beego"
)

type Api struct {
	BaseController
}

func (c *Api) AddCart() {
	id := c.GetSession("uid")
	productid := c.GetString("productid")
	amount := c.GetString("amount")

	err := models.Exec(`insert into shopcar (userid,productid,amount)
							values (?,?,?)`, id, productid, amount)
	if err != nil {
		beego.Debug(err)
		c.Data["json"] = 0
	} else {
		c.Data["json"] = 1
	}
	c.ServeJSON()
}

func (c *Api) UpdateCart() {
	id := c.GetSession("uid")
	productid := c.GetString("productid")
	amount := c.GetString("amount")

	err := models.Exec(`update shopcar set amount=? where userid=? and id=?`, amount, id, productid)
	if err != nil {
		beego.Debug(err)
		c.Data["json"] = 0
	} else {
		c.Data["json"] = 1
	}
	c.ServeJSON()
}

func (c *Api) GetCartList() {
	id := c.GetSession("uid")
	res, num := models.Querynum(`select s.*,p.name,p.price,p.images from shopcar s join products p on s.productid=p.id where userid=?`, id)
	if num == 0 {
		c.Data["json"] = 0
	} else {
		c.Data["json"] = res
	}
	c.ServeJSON()
}

func (c *Api) GetProductInfo() {
	productid := c.GetString("productid")
	res, num := models.Querynum(`select * from products where id=?`, productid)
	if num == 0 {
		c.Data["json"] = 0
	} else {
		c.Data["json"] = res
	}
	c.ServeJSON()
}
