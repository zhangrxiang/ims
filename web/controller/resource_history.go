package controller

import (
	"github.com/kataras/iris"
	"simple-ims/models"
	"simple-ims/utils"
	"strconv"
	"time"
)

func ResourceHistoryDelete(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	if err != nil {
		response(ctx, false, "资源ID不合法:"+err.Error(), nil)
		return
	}
	user := auth(ctx)
	if user == nil {
		return
	}
	rm := &models.ResourceModel{ID: id}
	rm, err = rm.First()
	if err != nil {
		response(ctx, false, "资源不存在:"+err.Error(), nil)
		return
	}
	rhm := &models.ResourceHistoryModel{ResourceID: rm.ID}
	rhm2, err2 := rhm.FindBy()
	if err2 != nil {
		response(ctx, false, "资源版本不存在:"+err2.Error(), nil)
		return
	}
	if len(rhm2) <= 1 {
		response(ctx, false, "资源版本数量小于一个", nil)
		return
	}
	err = rhm2[0].Delete()
	if err != nil {
		response(ctx, false, "删除当前版本失败:"+err.Error(), nil)
		return
	}
	rm.RHId = rhm2[1].ID
	rm, err = rm.Update()
	if err != nil {
		response(ctx, false, "更新版本失败:"+err.Error(), nil)
		return
	}
	response(ctx, true, "删除版本成功", nil)
}

func ResourceHistoryUpdate(ctx iris.Context) {
	resourceID, err := ctx.PostValueInt("resource_id")
	if err != nil {
		response(ctx, false, "资源ID非法", nil)
		return
	}

	rm := &models.ResourceModel{ID: resourceID}
	rm, err = rm.First()
	if err != nil {
		response(ctx, false, "资源不存在:"+err.Error(), nil)
		return
	}

	resourceHistoryModel := &models.ResourceHistoryModel{ID: rm.RHId}
	model, err := resourceHistoryModel.FirstBy()
	if err != nil {
		response(ctx, false, "当前版本不存在:"+err.Error(), nil)
		return
	}

	logStr := ctx.PostValue("log")
	file, info, err := ctx.FormFile("file")
	if file != nil {
		if err != nil {
			response(ctx, false, "获取上传文件失败:"+err.Error(), nil)
			return
		}
		_ = file.Close()

		uploadDir := "uploads/" + time.Now().Format("2006/01/")
		if !utils.Mkdir(uploadDir) {
			response(ctx, false, "创建文件夹失败", nil)
			return
		}

		resourceHistoryModel.Hash, err = utils.Md5File(file)
		if err != nil {
			response(ctx, false, "获取文件MD5失败:"+err.Error(), nil)
			return
		}
		resourceHistoryModel.Log = logStr
		resourceHistoryModel.File = info.Filename
		resourceHistoryModel.Path = uploadDir + utils.FileName(info.Filename, resourceHistoryModel.Version)
		err = utils.CopyFile(resourceHistoryModel.Path, file)
		if err != nil {
			response(ctx, false, "保存文件失败:"+err.Error(), nil)
			return
		}
	} else {
		response(ctx, false, "请上传文件", nil)
		return
	}

	model, err = resourceHistoryModel.Update()
	if err != nil {
		response(ctx, false, "更新当前资源版本失败:"+err.Error(), nil)
		return
	}

	resourceModel := &models.ResourceModel{
		ID:   resourceID,
		RHId: model.ID,
	}
	_, err = resourceModel.Update()
	if err != nil {
		response(ctx, false, "更新资源失败:"+err.Error(), nil)
		return
	}

	response(ctx, true, "更新当前资源版本成功", iris.Map{
		"resource": model,
	})
}

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
	historyModel, err := model.FindBy()
	if err != nil {
		response(ctx, false, "获取历史版本失败", nil)
		return
	}
	response(ctx, true, "", iris.Map{
		"resources": historyModel,
	})
}
