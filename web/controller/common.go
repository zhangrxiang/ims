package controller

import (
	"fmt"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"simple-ims/models"
	"simple-ims/utils"
)

type Message struct {
	Success bool        `json:"success"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}

//响应客户端数据
func response(ctx iris.Context, success bool, errMsg string, data interface{}) {
	if data == nil {
		data = []int{}
	}
	_, err := ctx.JSON(Message{
		Success: success,
		ErrMsg:  errMsg,
		Data:    data,
	})
	if !success {
		utils.Error(fmt.Sprintf("[message:%s],[data:%t]", errMsg, data))
	}
	if err != nil {
		utils.Error("输出json数据失败： ", err)
	}
}

func auth(ctx iris.Context) *models.UserModel {
	u := ctx.Values().Get("user")
	if u == nil {
		response(ctx, false, "请登录", nil)
		ctx.StopExecution()
		return nil
	}
	user := u.(map[string]interface{})
	return &models.UserModel{
		ID:       int(user["id"].(float64)),
		Username: user["username"].(string),
		Password: user["password"].(string),
		Role:     user["role"].(string),
		Phone:    user["phone"].(string),
		Mail:     user["mail"].(string),
	}
}

func authUser(ctx iris.Context) (*models.UserModel, error) {

	user := ctx.Values().Get("user")
	if user != nil {
		return user.(*models.UserModel), nil
	}
	token := ctx.Values().Get("jwt").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	model := &models.UserModel{
		ID: int(claims["userId"].(float64)),
	}
	userModel, err := model.FindByID()
	if err != nil {
		return nil, err
	}
	ctx.Values().Set("user", userModel)

	return userModel, nil
}
