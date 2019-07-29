package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"simple-ims/web/controllers/api/v1"
)

const V1 = "/api/v1"

var (
	app *iris.Application
)

func init() {
	app = iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./web/views", ".html"))
}

func routers() {

	app.Get("/", func(context context.Context) {
		_, _ = context.WriteString("PONG")
	})

	user := new(v1.UserController)
	app.Get(V1+"/login", user.Login)
	mvc.Configure(app.Party(V1+"/user"), func(app *mvc.Application) {
		//app.Router.Use(middleware.BasicAuth)
		app.Handle(user)
	})

}
func main() {
	routers()
	_ = app.Run(
		iris.Addr("localhost:8081"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
