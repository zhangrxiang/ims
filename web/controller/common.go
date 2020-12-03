package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"simple-ims/models"
	"simple-ims/utils"
)

type Message struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func login(user *models.UserModel) {
	lm := models.LogModel{
		UserId:  user.ID,
		Content: fmt.Sprintf("[ %s ] 登录资源管理系统", user.Username),
	}
	lm.Insert()
}

func log(ctx iris.Context, content string) {
	user := auth(ctx)
	lm := models.LogModel{
		UserId:  user.ID,
		Content: fmt.Sprintf("[ %s ] %s", user.Username, content),
	}
	lm.Insert()
}

//响应客户端数据
func response(ctx iris.Context, success bool, message string, data interface{}) {
	if data == nil {
		data = []int{}
	}
	_, err := ctx.JSON(Message{
		Success: success,
		Message: message,
		Data:    data,
	})
	if !success {
		utils.Error(fmt.Sprintf("[message:%s],[data:%v]", message, data))
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
