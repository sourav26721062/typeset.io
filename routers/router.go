package routers

import (
	"github.com/astaxie/beego"
	"typeset.io/controllers"
)

func init() {
	beego.Router("/api/post", &controllers.PostController{}, "get:GetAll")
	beego.Router("/api/post", &controllers.PostController{}, "post:Post")
	beego.Router("/api/post/:id", &controllers.PostController{}, "get:Get")
	beego.Router("/api/para/:pid/comment", &controllers.PostController{}, "post:PostComment")
	beego.Router("/api/para/:pid/comment", &controllers.PostController{}, "get:GetComment")

}
