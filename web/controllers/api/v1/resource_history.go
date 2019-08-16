package v1

import (
	"github.com/kataras/iris"
	"simple-ims/models"
	"strconv"
)

func ResourceHistoryLists(ctx iris.Context) {

	resourceIdStr := ctx.URLParam("resource_id")
	resourceId, err := strconv.Atoi(resourceIdStr)
	if err != nil {
		response(ctx, false, "资源ID非法", nil)
		return
	}
	model := models.ResourceHistoryModel{
		ResourceID: resourceId,
	}
	historyModel, err := model.FindByResourceId()
	if err != nil {
		response(ctx, false, "获取历史版本失败", nil)
		return
	}
	response(ctx, true, "", iris.Map{
		"resources": historyModel,
	})
}
