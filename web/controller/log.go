package controller

import (
	"github.com/kataras/iris"
	"simple-ims/models"
)

func LogList(ctx iris.Context) {
	user := auth(ctx)
	var (
		data []models.LogModel
		err  error
	)
	lm := &models.LogModel{}
	if user.Role == models.Admin {
		data, err = lm.Find()
	} else {
		data, err = lm.FindNot(1)
	}

	if err != nil {
		response(ctx, false, "", data)
		return
	}
	response(ctx, true, "", data)
}
