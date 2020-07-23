package routers

import (
	"day3/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/store", &controllers.MainController{}, "Get:Store")
	beego.Router("/product/list", &controllers.MainController{}, "Get:Products")
	beego.Router("/store/register", &controllers.Register{})
	beego.Router("/store/login", &controllers.Login{})
	beego.Router("/store/category", &controllers.MainController{}, "Get:Category")
	beego.Router("/store/cart", &controllers.MainController{}, "Get:Cart")
	beego.Router("/store/my", &controllers.MainController{}, "Get:My")
	beego.Router("/store/cart/add", &controllers.Api{}, "Post:AddCart")
	beego.Router("/store/product/detail", &controllers.MainController{}, "Get:Detail")
	beego.Router("/store/buy", &controllers.MainController{}, "Get:BuyP")
	beego.Router("/store/cart/getlist", &controllers.Api{}, "Post:GetCartList")
	beego.Router("/store/product/singleinfo", &controllers.Api{}, "Post:GetProductInfo")
}
