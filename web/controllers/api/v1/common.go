package v1

import (
	"github.com/kataras/iris"
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
	n, err := ctx.JSON(Message{
		Success: success,
		ErrMsg:  errMsg,
		Data:    data,
	})

	if err != nil {
		ctx.Application().Logger().Fatal("输出json数据失败,n:", n, err)
		return
	}
	return
}
