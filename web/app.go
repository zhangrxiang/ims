package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"log"
	"simple-ims/models"
	"simple-ims/utils"
	"simple-ims/web/controller"
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
	web.app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	tmpl := iris.HTML("./www", ".html")
	tmpl.Reload(true)
	web.app.RegisterView(tmpl)
	web.app.SPA(web.app.StaticHandler("./www", false, false))

	models.GetDBInstance()
	utils.LoadLogInit()

	//ping
	web.app.Get("/ping", func(context context.Context) {
		_, _ = context.WriteString("PONG")
	})

	tmp := web.app.Party(v1Api + "/tmp")
	{
		tmp.Get("/lists", controller.TmpUpLists)
		tmp.Post("/upload", controller.TmpUpload)
		tmp.Get("/download", controller.TmpDownload)
	}

	//用户
	web.app.Get(v1Api+"/user/login", controller.UserLogin)
	user := web.app.Party(v1Api + "/user")
	user.Use(middleware.JWT)
	user.Use(middleware.Auth)
	{
		user.Get("/lists", controller.UserLists)
		user.Post("/register", controller.UserRegister)
		user.Get("/delete", controller.UserDelete)
		user.Post("/update", controller.UserUpdate)
	}

	//资源分类
	resourceType := web.app.Party(v1Api + "/resource-type")
	resourceType.Use(middleware.JWT)
	resourceType.Use(middleware.Auth)
	{
		resourceType.Post("/add", controller.ResourceTypeAdd)
		resourceType.Get("/lists", controller.ResourceTypeLists)
		resourceType.Post("/update", controller.ResourceTypeUpdate)
		resourceType.Get("/delete", controller.ResourceTypeDelete)
	}

	//资源
	resource := web.app.Party(v1Api + "/resource")
	resource.Use(middleware.JWT)
	resource.Use(middleware.Auth)
	{
		resource.Post("/add", controller.ResourceAdd)
		resource.Get("/lists", controller.ResourceLists)
		resource.Get("/delete", controller.ResourceDelete)
		resource.Post("/update", controller.ResourceUpdate)
		resource.Post("/upgrade", controller.ResourceUpgrade)
		resource.Get("/download", controller.ResourceDownload)
		resource.Get("/group-lists", controller.ResourceGroupLists)
	}

	//历史版本
	resourceHistory := web.app.Party(v1Api + "/resource-history")
	resourceHistory.Use(middleware.JWT)
	resourceHistory.Use(middleware.Auth)
	{
		resourceHistory.Get("/lists", controller.ResourceHistoryLists)
	}

	//项目
	project := web.app.Party(v1Api + "/project")
	project.Use(middleware.JWT)
	project.Use(middleware.Auth)
	{
		project.Post("/add", controller.ProjectAdd)
		project.Post("/update", controller.ProjectUpdate)
		project.Post("/upgrade", controller.ProjectUpgrade)
		project.Get("/detail", controller.ProjectDetail)
		project.Get("/lists", controller.ProjectLists)
		project.Get("/delete", controller.ProjectDelete)
		project.Get("/download", controller.ProjectDownload)
	}

	projectHistory := web.app.Party(v1Api + "/project-history")
	projectHistory.Use(middleware.JWT)
	projectHistory.Use(middleware.Auth)
	{
		projectHistory.Get("/lists", controller.ProjectHistoryLists)
	}

	go func() {
		utils.Info("web start...")
		err := web.app.Run(
			iris.Addr(":5050"),
			iris.WithoutServerError(iris.ErrServerClosed),
			iris.WithOptimizations,
		)

		if err != nil {
			log.Fatalln("app run error ", err)
		}
	}()
}
