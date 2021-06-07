package web

import (
	"github.com/kataras/iris/v12"
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
	utils.InitLog()
	web.app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))
	web.app.HandleDir("/", iris.Dir("./www"), iris.DirOptions{
		ShowList:  true,
		IndexName: "index.html",
		SPA:       true,
	})
	models.GetDBInstance()
	//ping
	web.app.Get("/ping", func(context iris.Context) {
		_, _ = context.WriteString("PONG")
	})

	tmp := web.app.Party(v1Api + "/tmp")
	{
		tmp.Get("/lists", controller.TmpUpLists)
		tmp.Post("/upload", controller.TmpUpload)
		tmp.Get("/download", controller.TmpDownload)
	}

	v := web.app.Party(v1Api + "/version")
	{
		v.Get("/{name}/lists", controller.VersionLists)
		v.Get("/download/{id}", controller.VersionDownload)
	}

	//用户
	web.app.Get(v1Api+"/user/login", controller.UserLogin)
	user := web.app.Party(v1Api + "/user")
	user.Use(middleware.JWT, middleware.Auth)
	{
		user.Get("/lists", controller.UserLists)
		user.Post("/register", controller.UserRegister)
		user.Get("/delete", controller.UserDelete)
		user.Post("/update", controller.UserUpdate)
	}

	//资源分类
	resourceType := web.app.Party(v1Api + "/resource-type")
	resourceType.Use(middleware.JWT, middleware.Auth)
	resourceType.Use()
	{
		resourceType.Post("/add", controller.ResourceTypeAdd)
		resourceType.Get("/lists", controller.ResourceTypeLists)
		resourceType.Post("/update", controller.ResourceTypeUpdate)
		resourceType.Get("/delete", controller.ResourceTypeDelete)
	}

	//资源
	resource := web.app.Party(v1Api + "/resource")
	resource.Use(middleware.JWT, middleware.Auth)
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
	resourceHistory.Use(middleware.JWT, middleware.Auth)
	{
		resourceHistory.Get("/delete", controller.ResourceHistoryDelete)
		resourceHistory.Get("/rollback", controller.ResourceHistoryRollback)
		resourceHistory.Post("/update", controller.ResourceHistoryUpdate)
		resourceHistory.Get("/lists", controller.ResourceHistoryLists)
	}

	//项目
	project := web.app.Party(v1Api + "/project")
	project.Use(middleware.JWT, middleware.Auth)
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
	projectHistory.Use(middleware.JWT, middleware.Auth)
	{
		projectHistory.Get("/lists", controller.ProjectHistoryLists)
	}

	remote := web.app.Party(v1Api + "/remote")
	remote.Use(middleware.JWT, middleware.Auth)
	{
		c := new(controller.Remote)
		remote.Get("/exec", c.Exec)
		remote.Get("/list", c.List)
		remote.Get("/download", c.Download)
	}

	web.app.Get(v1Api+"/log/lists", middleware.JWT, middleware.Auth, controller.LogList)
	web.app.Get(v1Api+"/info", middleware.JWT, middleware.Auth, controller.Info)

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
