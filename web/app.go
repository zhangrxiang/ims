package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"log"
	"simple-ims/models"
	"simple-ims/web/controllers/api/v1"
	"simple-ims/web/middleware"
	"sync"
)

type Web struct {
	app *iris.Application
}

var (
	webOnce sync.Once
	w       *Web
)

const v1Api = "/api/v1"

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
	web.app.Logger().SetOutput(newLogFile())
	web.app.Logger().SetLevel("debug")
	web.app.Use(logger.New())
	web.app.SPA(web.app.StaticHandler("./www", false, false))
	web.app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	//ping
	web.app.Get("/ping", func(context context.Context) {
		_, _ = context.WriteString("PONG")
	})

	models.GetDBInstance()

	//用户
	web.app.Get(v1Api+"/user/login", v1.UserLogin)
	user := web.app.Party(v1Api + "/user")
	user.Use(middleware.JWT)
	{
		user.Get("/lists", v1.UserLists)
		user.Post("/register", v1.UserRegister)
		user.Get("/delete", v1.UserDelete)
		user.Post("/update", v1.UserUpdate)
	}

	//资源分类
	resourceType := web.app.Party(v1Api + "/resource-type")
	resourceType.Use(middleware.JWT)
	{
		resourceType.Post("/add", v1.ResourceTypeAdd)
		resourceType.Get("/lists", v1.ResourceTypeLists)
		resourceType.Post("/update", v1.ResourceTypeUpdate)
		resourceType.Get("/delete", v1.ResourceTypeDelete)
	}

	//资源
	resource := web.app.Party(v1Api + "/resource")
	resource.Use(middleware.JWT)
	{
		resource.Post("/add", v1.ResourceAdd)
		resource.Get("/lists", v1.ResourceLists)
		resource.Get("/delete", v1.ResourceDelete)
		resource.Post("/update", v1.ResourceUpdate)
	}
	web.app.Get(v1Api+"/resource/download", v1.ResourceDownload)
	web.app.Get(v1Api+"/resource/group-lists", v1.ResourceGroupLists)

	//web.app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
	//	_, _ = ctx.JSON(iris.Map{
	//		"success": false,
	//		"err_msg": iris.StatusNotFound,
	//		"data":    []int{},
	//	})
	//})

	go func() {
		err := web.app.Run(
			iris.Addr(":8081"),
			iris.WithoutServerError(iris.ErrServerClosed),
			iris.WithOptimizations,
		)

		if err != nil {
			log.Fatalln("app run error ", err)
		}
	}()
}
