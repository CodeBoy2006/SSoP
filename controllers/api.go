package controllers

import (
	"day3/models"
	"strconv"

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

	res, num := models.Querynum(`select amount from shopcar where userid=? and productid=?`, id, productid)
	if num == 0 {
		err := models.Exec(`insert into shopcar (userid,productid,amount)
							values (?,?,?)`, id, productid, amount)
		if err != nil {
			beego.Debug(err)
			c.Data["json"] = 0

		} else {
			c.Data["json"] = 1
		}
	} else {
		ramount, _ := strconv.Atoi(res[0]["amount"].(string))
		beego.Debug(ramount)
		err := models.Exec(`update shopcar set amount=? where userid=? and productid=?`, (ramount + 1), id, productid)
		if err != nil {
			beego.Debug(err)
			c.Data["json"] = 0
		} else {
			c.Data["json"] = 1
		}
	}

	c.ServeJSON()
}

func (c *Api) UpdateCart() {
	mode := c.GetString("mode")
	if mode == "uv" {
		id := c.GetSession("uid")
		productid := c.GetString("productid")
		check := c.GetString("check")
		amount := c.GetString("amount")

		err := models.Exec(`update shopcar set amount=?, ischeck=? where userid=? and id=?`, amount, check, id, productid)
		if err != nil {
			beego.Debug(err)
			c.Data["json"] = 0
		} else {
			c.Data["json"] = 1
		}
	}
	if mode == "del" {
		uid := c.GetSession("uid")
		id := c.GetString("id")
		err := models.Exec(`delete from shopcar where id=? and userid=?`, id, uid)
		if err != nil {
			beego.Debug(err)
			c.Data["json"] = 0
		} else {
			c.Data["json"] = 1
		}
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

func (c *Api) GetUserInfoBySession() {
	userid := c.GetSession("uid")
	res, num := models.Querynum(`select username,email,phone from users where id=?`, userid)
	if num == 0 {
		c.Data["json"] = 0
	} else {
		c.Data["json"] = res[0]
	}
	c.ServeJSON()
}
