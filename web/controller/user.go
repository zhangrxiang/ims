package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"regexp"
	"simple-ims/models"
	"simple-ims/utils"
	"simple-ims/web/middleware"
	"strings"
	"time"
)

//用户列表
func UserLists(ctx iris.Context) {
	var (
		user  models.UserModel
		users []models.UserModel
		err   error
	)
	current := auth(ctx)
	if current.Role == models.Admin {
	} else {
		user.ID = current.ID
	}
	users, err = user.Find()
	if err != nil {
		response(ctx, false, "无用户:"+err.Error(), nil)
		return
	}
	for k := range users {
		users[k].Password = utils.Decode(users[k].Password)
	}
	response(ctx, true, "", iris.Map{
		"users":     users,
		"timestamp": time.Now().Unix(),
	})
}

//用户登陆
func UserLogin(ctx iris.Context) {
	username := ctx.URLParamDefault("username", "")
	password := ctx.URLParamDefault("password", "")
	if username == "" || password == "" {
		response(ctx, false, "用户名或密码不能为空", nil)
		return
	}
	u := &models.UserModel{
		Username: username,
		Password: utils.Encode(password),
	}
	u, err := u.First()
	if err != nil {
		response(ctx, false, "用户名或密码错误", nil)
		return
	}

	u.Password = strings.Repeat("*", len(u.Password))
	token, err := middleware.GenerateToken(u)

	if err != nil {
		response(ctx, false, "生成token失败"+err.Error(), nil)
		return
	}
	login(u)
	response(ctx, true, "", iris.Map{
		"user":       u,
		"token":      token,
		"token_type": "Bearer",
		"timestamp":  time.Now().Unix(),
	})
}

//用户注册
func UserRegister(ctx iris.Context) {
	if auth(ctx).Role != models.Admin {
		response(ctx, false, "没有注册权限,禁止操作", nil)
		return
	}
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

	um := &models.UserModel{
		Username: username,
		Password: utils.Encode(password),
		Mail:     mail,
		Phone:    phone,
		Role:     role,
	}

	if err := um.Insert(); err != nil {
		response(ctx, false, "注册用户失败:"+err.Error(), nil)
		return
	}
	log(ctx, fmt.Sprintf("注册用户:[ %s ], 密码[ %s ], 角色[ %s ]", username, password, role))
	response(ctx, true, "注册用户成功", iris.Map{
		"user": um,
	})
}

//用户删除
func UserDelete(ctx iris.Context) {
	if auth(ctx).Role != models.Admin {
		response(ctx, false, "没有删除权限,禁止操作", nil)
		return
	}
	ids := utils.StrToIntSlice(ctx.FormValue("id"), ",")
	if ids == nil {
		response(ctx, false, "用户ID非法", nil)
		return
	}
	for _, id := range ids {
		um := &models.UserModel{ID: id}
		um, err := um.First()
		if err != nil {
			continue
		}
		if err := um.Delete(); err == nil {
			log(ctx, fmt.Sprintf("删除用户:[ %s ], 密码[ %s ], 角色[ %s ]", um.Username, utils.Decode(um.Password), um.Role))
		}
	}
	response(ctx, true, "删除用户成功", nil)
}

//用户更新
func UserUpdate(ctx iris.Context) {
	id, err := ctx.PostValueInt("id")
	if err != nil {
		response(ctx, false, "用户ID非法", nil)
		return
	}
	user := auth(ctx)
	if user.Role != models.Admin && user.ID != id {
		response(ctx, false, "没有修改信息权限,禁止操作", nil)
		return
	}
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
	um := &models.UserModel{
		ID:       id,
		Username: username,
		Password: utils.Encode(password),
		Mail:     mail,
		Phone:    phone,
		Role:     role,
	}

	if err := um.Update(); err != nil {
		response(ctx, false, "修改用户失败:"+err.Error(), nil)
		return
	}
	log(ctx, "更新用户信息")
	response(ctx, true, "修改用户成功", nil)
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
