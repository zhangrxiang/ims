package v1

import (
	"github.com/kataras/iris"
	"simple-ims/models"
	"time"
)

//用户列表
func UserList(ctx iris.Context) {

	users, err := (&models.UserModel{}).All()
	if err != nil {
		response(ctx, false, "无用户:"+err.Error(), nil)
		return
	}

	response(ctx, true, "", iris.Map{
		"users":     users,
		"timestamp": time.Now().Unix(),
	})
	return

}

//用户登陆
func UserLogin(ctx iris.Context) {

	username := ctx.URLParamDefault("username", "")
	password := ctx.URLParamDefault("password", "")

	if username == "" || password == "" {
		response(ctx, false, "用户名或密码不能为空", nil)
		return
	}

	user := models.UserModel{
		Username: username,
		Password: password,
	}

	model, err := user.Find()

	if err != nil {
		response(ctx, false, "无此用户:"+err.Error(), nil)
		return
	}
	response(ctx, true, "", iris.Map{
		"user":      model,
		"timestamp": time.Now().Unix(),
	})
	return
}
