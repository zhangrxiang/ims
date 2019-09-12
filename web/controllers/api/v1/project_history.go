package v1

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
	phm := models.ProjectHistoryModel{
		ProjectId: projectId,
	}
	phms, err := phm.FindBy()
	if err != nil {
		response(ctx, false, "获取项目历史版本失败", nil)
		return
	}
	pm := &models.ProjectModel{
		ID: projectId,
	}
	pm, err = pm.FirstBy()
	if err != nil {
		response(ctx, false, "获取项目失败", nil)
		return
	}
	type project struct {
		Name string `json:"name"`
		Desc string `json:"desc"`
		models.ProjectHistoryModel
	}
	var data []project
	for _, v := range phms {
		data = append(data, project{pm.Name, pm.Desc, v})
	}
	response(ctx, true, "", data)
}
