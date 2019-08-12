package v1

import (
	"github.com/kataras/iris"
	"regexp"
	"simple-ims/models"
	"simple-ims/utils"
	"strconv"
	"time"
)

//用户列表
func UserLists(ctx iris.Context) {

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

//用户注册
func UserRegister(ctx iris.Context) {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	role := ctx.FormValue("role")
	mail := ctx.FormValue("mail")
	phone := ctx.FormValue("phone")

	if username == "" || password == "" || role == "" {
		response(ctx, false, "请输入用户名,密码,选择角色", nil)
		return
	}

	if mail != "" && !checkMail(mail) {
		response(ctx, false, "请输入合法的邮箱", nil)
		return
	}

	if phone != "" && !checkPhone(phone) {
		response(ctx, false, "请输入合法的手机号", nil)
		return
	}

	userModel := &models.UserModel{
		Username: username,
		Password: password,
		Mail:     mail,
		Phone:    phone,
		Role:     role,
	}

	model, err := userModel.Insert()
	if err != nil {
		response(ctx, false, "注册用户失败:"+err.Error(), nil)
		return
	}

	response(ctx, true, "注册用户成功", iris.Map{
		"user": model,
	})

}

//用户删除
func UserDelete(ctx iris.Context) {
	id := ctx.FormValue("id")
	ids := utils.StrToIntAlice(id, ",")
	if ids == nil {
		response(ctx, false, "用户ID非法", nil)
		return
	}
	userModel := &models.UserModel{}
	_, err := userModel.Delete(ids)

	if err != nil {
		response(ctx, false, "删除用户失败:"+err.Error(), nil)
		return
	}
	response(ctx, false, "删除用户成功", nil)
}

//用户更新
func UserUpdate(ctx iris.Context) {
	idStr := ctx.FormValue("id")
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	role := ctx.FormValue("role")
	mail := ctx.FormValue("mail")
	phone := ctx.FormValue("phone")

	if idStr == "" {
		response(ctx, false, "用户ID非法", nil)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response(ctx, false, "用户ID非法", nil)
		return
	}
	if username == "" || password == "" || role == "" {
		response(ctx, false, "请输入用户名,密码,选择角色", nil)
		return
	}

	if mail != "" && !checkMail(mail) {
		response(ctx, false, "请输入合法的邮箱", nil)
		return
	}

	if phone != "" && !checkPhone(phone) {
		response(ctx, false, "请输入合法的手机号", nil)
		return
	}

	userModel := &models.UserModel{
		ID:       id,
		Username: username,
		Password: password,
		Mail:     mail,
		Phone:    phone,
		Role:     role,
	}

	model, err := userModel.Update()
	if err != nil {
		response(ctx, false, "修改用户失败:"+err.Error(), nil)
		return
	}

	response(ctx, true, "修改用户成功", iris.Map{
		"user": model,
	})
}

//验证邮箱
func checkMail(mail string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(mail)
}

//验证手机号
func checkPhone(phone string) bool {
	pattern := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(phone)
}
