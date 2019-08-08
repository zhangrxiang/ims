package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"log"
	v1 "simple-ims/web/controllers/api/v1"
	"sync"
)

type Web struct {
	app *iris.Application
}

var (
	webOnce sync.Once
	w       *Web
)

func NewOnceWeb() *Web {
	webOnce.Do(func() {
		w := &Web{
			app: iris.New(),
		}
		w.Init()
	})
	return w
}

func (web *Web) Init() {

	web.app.Logger().SetLevel("debug")
	web.app.Use(logger.New())
	web.app.RegisterView(iris.HTML("./web/views", ".html"))
	web.app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	web.app.Get("/", func(context context.Context) {
		_, _ = context.WriteString("PONG")
	})

	web.app.Get("/api/v1/user/login", v1.UserLogin)
	web.app.Get("/api/v1/user/lists", v1.UserList)

	web.app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		_, _ = ctx.JSON(iris.Map{
			"success": false,
			"err_msg": iris.StatusNotFound,
			"data":    []int{},
		})
	})

	err := web.app.Run(
		iris.Addr("localhost:8081"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

	if err != nil {
		log.Fatalln("app run error ", err)
	}
}
