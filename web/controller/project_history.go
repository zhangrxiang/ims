package controller

import (
	"github.com/kataras/iris"
	"simple-ims/models"
)

func ProjectHistoryLists(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("project_id")
	if err != nil {
		response(ctx, false, "项目ID非法", nil)
		return
	}
	phm := &models.ProjectHistoryModel{ProjectId: projectId}
	list, err := phm.Find()
	if err != nil {
		response(ctx, false, "获取项目历史版本失败", nil)
		return
	}
	pm := &models.ProjectModel{Id: projectId}
	pm, err = pm.First()
	if err != nil {
		response(ctx, false, "获取项目失败", nil)
		return
	}
	type item struct {
		Name string `json:"name"`
		Desc string `json:"desc"`
		models.ProjectHistoryModel
	}
	var data []item
	for _, v := range list {
		data = append(data, item{pm.Name, pm.Desc, v})
	}
	response(ctx, true, "", data)
}
