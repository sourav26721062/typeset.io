package controllers

import (
	"github.com/astaxie/beego"
)

//ErrorController - struct defination
type ErrorController struct {
	beego.Controller
}

//Error404 - return error 404 api not found
func (_error *ErrorController) Error404() {
	_error.Data["json"] = "api not found"
	_error.Ctx.ResponseWriter.WriteHeader(404)
	_error.ServeJSON()
}

//Error500 - return error 500 for internal server error
func (_error *ErrorController) Error500() {
	_error.Data["content"] = "internal server error"
	_error.Ctx.ResponseWriter.WriteHeader(505)
	_error.ServeJSON()
}

//ErrorDb - return error 403 if Db is down
func (_error *ErrorController) ErrorDb() {
	_error.Data["content"] = "database is now down"
	_error.Ctx.ResponseWriter.WriteHeader(403)
	_error.ServeJSON()
}
