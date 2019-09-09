package v1

import (
	"github.com/kataras/iris"
	"simple-ims/models"
)

//添加项目
func ProjectAdd(ctx iris.Context) {
	name := ctx.FormValue("name")
	desc := ctx.FormValue("desc")
	if name == "" || desc == "" {
		response(ctx, false, "请输入项目名称和简介", nil)
		return
	}
	pm := &models.ProjectModel{
		Name:   name,
		Desc:   desc,
		UserId: auth(ctx).ID,
	}
	model, err := pm.Insert()
	if err != nil {
		response(ctx, false, "保存项目失败:"+err.Error(), nil)
		return
	}
	response(ctx, false, "保存项目成功", model)
}
