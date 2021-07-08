package controller

import (
	"github.com/kataras/iris/v12"
	"simple-ims/models"
)

func Info(ctx iris.Context) {
	response(ctx, true, "", models.GetVersion())
}
