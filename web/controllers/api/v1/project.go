package v1

import (
	"github.com/kataras/iris"
	"simple-ims/models"
	"simple-ims/utils"
)

//项目列表
func ProjectLists(ctx iris.Context) {
	pm := models.ProjectModel{}
	model, err := pm.FindBy()
	if err != nil {
		response(ctx, true, "获取项目列表失败:"+err.Error(), nil)
		return
	}
	response(ctx, true, "获取项目列表成功", model)
}

//添加项目版本
func ProjectUpgrade(ctx iris.Context) {
	projectId, err := ctx.PostValueInt("project_id")
	if err != nil {
		response(ctx, false, "项目ID非法", nil)
		return
	}
	version := ctx.FormValue("version")
	resourceIds := ctx.FormValue("resource_ids")
	log := ctx.FormValue("log")
	if version == "" || resourceIds == "" || log == "" {
		response(ctx, false, "请输入版本号和日志,选择对应资源", nil)
		return
	}
	phm := models.ProjectHistoryModel{
		ProjectId:   projectId,
		Version:     version,
		Log:         log,
		ResourceIds: resourceIds,
		Hash:        utils.Md5Str(string(projectId) + version + log + resourceIds),
	}
	model, err := phm.Insert()
	if err != nil {
		response(ctx, false, "保存项目版本失败:"+err.Error(), nil)
		return
	}
	pm := models.ProjectModel{
		ID:   projectId,
		PHId: model.ID,
	}
	_, err = pm.Update()
	if err != nil {
		response(ctx, false, "更新项目失败:"+err.Error(), model)
		return
	}
	response(ctx, false, "保存项目版本成功", model)
}

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
