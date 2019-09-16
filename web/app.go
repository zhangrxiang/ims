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
	web.app.Logger().SetLevel("warn")
	web.app.Use(logger.New())
	web.app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	tmpl := iris.HTML("./www", ".html")
	tmpl.Reload(true)
	web.app.RegisterView(tmpl)
	web.app.SPA(web.app.StaticHandler("./www", false, false))

	models.GetDBInstance()

	//ping
	web.app.Get("/ping", func(context context.Context) {
		_, _ = context.WriteString("PONG")
	})
	//用户
	web.app.Get(v1Api+"/user/login", v1.UserLogin)
	user := web.app.Party(v1Api + "/user")
	user.Use(middleware.JWT)
	user.Use(middleware.Auth)
	{
		user.Get("/lists", v1.UserLists)
		user.Post("/register", v1.UserRegister)
		user.Get("/delete", v1.UserDelete)
		user.Post("/update", v1.UserUpdate)
	}

	//资源分类
	resourceType := web.app.Party(v1Api + "/resource-type")
	resourceType.Use(middleware.JWT)
	resourceType.Use(middleware.Auth)
	{
		resourceType.Post("/add", v1.ResourceTypeAdd)
		resourceType.Get("/lists", v1.ResourceTypeLists)
		resourceType.Post("/update", v1.ResourceTypeUpdate)
		resourceType.Get("/delete", v1.ResourceTypeDelete)
	}

	//资源
	resource := web.app.Party(v1Api + "/resource")
	resource.Use(middleware.JWT)
	resource.Use(middleware.Auth)
	{
		resource.Post("/add", v1.ResourceAdd)
		resource.Get("/lists", v1.ResourceLists)
		resource.Get("/delete", v1.ResourceDelete)
		resource.Post("/update", v1.ResourceUpdate)
		resource.Post("/upgrade", v1.ResourceUpgrade)
		resource.Get("/download", v1.ResourceDownload)
		resource.Get("/group-lists", v1.ResourceGroupLists)
	}

	//历史版本
	resourceHistory := web.app.Party(v1Api + "/resource-history")
	resourceHistory.Use(middleware.JWT)
	resourceHistory.Use(middleware.Auth)
	{
		resourceHistory.Get("/lists", v1.ResourceHistoryLists)
	}

	//项目
	project := web.app.Party(v1Api + "/project")
	project.Use(middleware.JWT)
	project.Use(middleware.Auth)
	{
		project.Post("/add", v1.ProjectAdd)
		project.Post("/upgrade", v1.ProjectUpgrade)
		project.Get("/lists", v1.ProjectLists)
		project.Get("/delete", v1.ProjectDelete)
		//project.Post("/update", v1.ResourceUpdate)
		project.Get("/download", v1.ProjectDownload)
		//project.Get("/group-lists", v1.ResourceGroupLists)
	}

	projectHistory := web.app.Party(v1Api + "/project-history")
	projectHistory.Use(middleware.JWT)
	projectHistory.Use(middleware.Auth)
	{
		projectHistory.Get("/lists", v1.ProjectHistoryLists)
	}

	//web.app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
	//	_, _ = ctx.JSON(iris.Map{
	//		"success": false,
	//		"err_msg": iris.StatusNotFound,
	//		"data":    []int{},
	//	})
	//})

	go func() {
		err := web.app.Run(
			iris.Addr(":80"),
			iris.WithoutServerError(iris.ErrServerClosed),
			iris.WithOptimizations,
		)

		if err != nil {
			log.Fatalln("app run error ", err)
		}
	}()
}
