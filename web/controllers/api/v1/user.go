package v1

import (
	"github.com/kataras/iris"
	"simple-ims/models"
	"simple-ims/services"
)

type UserController struct {
	service services.UserService
}

func (u *UserController) GetBy(id int) (*models.UserModel, error) {
	return u.service.Find(int(id))
}

func (u *UserController) Login(ctx iris.Context) {
	username := ctx.URLParam("username")
	password := ctx.URLParam("password")
	model, err := u.service.Add(models.UserModel{
		Username: username,
		Password: password,
	})

	if err != nil {
		_, _ = ctx.JSON(iris.Map{
			"message": "注册失败",
			"status":  false,
		})
		return
	}
	_, _ = ctx.JSON(iris.Map{
		"message": "注册成功",
		"status":  true,
		"data":    model,
	})
}
