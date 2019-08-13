package v1

import (
	"fmt"
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
	if success {
		ctx.Application().Logger().Info(fmt.Sprintf("[success:%t],[err_msg:%s],[data:%t]", success, errMsg, data))
	} else {
		ctx.Application().Logger().Warn(fmt.Sprintf("[success:%t],[err_msg:%s],[data:%t]", success, errMsg, data))
	}
	if err != nil {
		ctx.Application().Logger().Warn("输出json数据失败,n:", n, err)
		return
	}
	return
}
