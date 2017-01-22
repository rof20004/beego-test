package routers

import (
	"mytest/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// Main page
	beego.Router("/", &controllers.MainController{}, "get:Get")

	// Usuario pages
	beego.Router("/usuario/index", &controllers.UsuarioController{}, "get:Index")
	beego.Router("/usuario/create", &controllers.UsuarioController{}, "get:Create")
	beego.Router("/usuario/edit/:id", &controllers.UsuarioController{}, "get:Edit")
	beego.Router("/usuario/view/:id", &controllers.UsuarioController{}, "get:View")

	// Usuario endpoints
	beego.Router("/usuario/list", &controllers.UsuarioController{}, "get:GetAll")
	beego.Router("/usuario/get/:id", &controllers.UsuarioController{}, "get:GetOne")
	beego.Router("/usuario/edit/:id", &controllers.UsuarioController{}, "put:Put")
	beego.Router("/usuario/save", &controllers.UsuarioController{}, "post:Post")
	beego.Router("/usuario/delete/:id", &controllers.UsuarioController{}, "delete:Delete")
}
