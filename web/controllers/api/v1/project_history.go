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
	historyModel, err := phm.FindBy()
	if err != nil {
		response(ctx, false, "获取项目历史版本失败", nil)
		return
	}
	response(ctx, true, "", historyModel)
}
