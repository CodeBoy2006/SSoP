package main

import (
	_ "day3/models"
	_ "day3/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

