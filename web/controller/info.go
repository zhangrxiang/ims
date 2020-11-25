package controller

import (
	"github.com/kataras/iris"
	"simple-ims/models"
)

func Info(ctx iris.Context) {
	response(ctx, true, "", models.Info)
}
