package main

import (
	"github.com/astaxie/beego"
	"typeset.io/controllers"
	_ "typeset.io/routers"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
