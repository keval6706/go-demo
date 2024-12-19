package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"go.demo/controller"
)

func MemberRouter(app *mvc.Application) {
	// app.Register(...)
	// app.Router.Use/UseGlobal/Done(...)
	app.Handle(new(controller.MemberController))
}
