package v1

import (
	"github.com/kataras/iris"
	"log"
	"os"
	"simple-ims/models"
	"simple-ims/utils"
	"strconv"
	"time"
)

//添加资源
func ResourceAdd(ctx iris.Context) {
	name := ctx.FormValue("name")
	_type := ctx.FormValue("type")
	version := ctx.FormValue("version")
	desc := ctx.FormValue("desc")
	logStr := ctx.FormValue("log")

	user, err := authUser(ctx)
	if err != nil {
		response(ctx, false, "请登录", nil)
		return
	}
	if user.Role != "admin" {
		response(ctx, false, "只有管理员才能添加资源", nil)
		return
	}

	if name == "" || _type == "" {
		response(ctx, false, "请输入资源名称,选择资源类型", nil)
		return
	}

	t, err := strconv.Atoi(_type)
	if err != nil {
		response(ctx, false, "资源类型不存在:"+err.Error(), nil)
		return
	}
	var resourceModel = &models.ResourceModel{
		Name:    name,
		Type:    t,
		Version: version,
		Desc:    desc,
		UserId:  user.ID,
	}

	file, info, err := ctx.FormFile("file")
	if file != nil {
		if err != nil {
			response(ctx, false, "获取上传文件失败:"+err.Error(), nil)
			return
		}

		defer file.Close()

		uploadDir := "uploads/" + time.Now().Format("2006/01/")
		if !utils.Mkdir(uploadDir) {
			response(ctx, false, "创建文件夹失败", nil)
			return
		}

		resourceModel.Hash, err = utils.Md5File(file)
		if err != nil {
			response(ctx, false, "获取文件MD5失败:"+err.Error(), nil)
			return
		}
		model, err := resourceModel.FindByHash(resourceModel.Hash)
		if model != nil {
			response(ctx, true, "相同的文件已存在:", iris.Map{
				"resource": model,
			})
			return
		}
		resourceModel.File = info.Filename
		resourceModel.Path = uploadDir + utils.FileName(info.Filename, version)
		err = utils.CopyFile(resourceModel.Path, file)
		if err != nil {
			response(ctx, false, "保存文件失败:"+err.Error(), nil)
			return
		}
	}

	resourceModel.CreateAt = time.Now()
	model, err := resourceModel.Insert()
	if err != nil {
		response(ctx, false, "添加资源失败:"+err.Error(), nil)
		return
	}

	resourceHistoryModel := &models.ResourceHistoryModel{
		ResourceID: model.ID,
		File:       model.File,
		Version:    model.Version,
		Path:       model.Path,
		Hash:       model.Hash,
		CreateAt:   model.CreateAt,
		Log:        logStr,
	}
	_, err = resourceHistoryModel.Insert()
	if err != nil {
		response(ctx, false, "保存历史资源版本失败:"+err.Error(), nil)
		return
	}

	response(ctx, true, "保存文件成功", iris.Map{
		"resource": model,
	})
}

//删除资源
func ResourceDelete(ctx iris.Context) {
	idsStr := ctx.FormValue("id")

	ids := utils.StrToIntAlice(idsStr, ",")

	if ids == nil {
		response(ctx, false, "资源ID不合法", nil)
		return
	}

	resourceModel := &models.ResourceModel{}
	resources, err := resourceModel.FindByIds(ids)
	if err != nil {
		response(ctx, false, "查找要删除的资源失败:"+err.Error(), nil)
		return
	}
	_, err = resourceModel.DeleteByIds(ids)
	if err != nil {
		response(ctx, false, "删除资源失败:"+err.Error(), nil)
		return
	}
	go func(resources *[]models.ResourceModel) {
		for _, resource := range *resources {
			err := os.Remove(resource.Path)
			if err != nil {
				log.Println("remove file", resource.Name, err)
			} else {
				log.Println("remove file", resource.Name)
			}
		}
	}(resources)

	response(ctx, true, "", nil)
	return
}

//更新资源
func ResourceUpdate(ctx iris.Context) {
	idStr := ctx.FormValue("id")
	name := ctx.FormValue("name")
	typeStr := ctx.FormValue("type")
	version := ctx.FormValue("version")
	desc := ctx.FormValue("desc")
	logStr := ctx.FormValue("log")

	if name == "" || typeStr == "" {
		response(ctx, false, "请输入资源名称,选择资源类型", nil)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response(ctx, false, "资源ID格式不合法:"+err.Error(), nil)
		return
	}
	t, err := strconv.Atoi(typeStr)
	if err != nil {
		response(ctx, false, "资源类型格式不合法:"+err.Error(), nil)
		return
	}
	var resourceModel = &models.ResourceModel{
		ID:      id,
		Name:    name,
		Type:    t,
		Version: version,
		Desc:    desc,
	}

	file, info, err := ctx.FormFile("file")
	if file != nil {
		if err != nil {
			response(ctx, false, "获取上传文件失败:"+err.Error(), nil)
			return
		}

		defer file.Close()

		uploadDir := "uploads/" + time.Now().Format("2006/01/")
		if !utils.Mkdir(uploadDir) {
			response(ctx, false, "创建文件夹失败", nil)
			return
		}

		resourceModel.Hash, err = utils.Md5File(file)
		if err != nil {
			response(ctx, false, "获取文件MD5失败:"+err.Error(), nil)
			return
		}
		model, err := resourceModel.FindByHash(resourceModel.Hash)
		if model != nil {
			response(ctx, true, "相同的文件已存在:", iris.Map{
				"resource": model,
			})
			return
		}
		resourceModel.File = info.Filename
		resourceModel.Path = uploadDir + utils.FileName(info.Filename, version)
		err = utils.CopyFile(resourceModel.Path, file)
		if err != nil {
			response(ctx, false, "保存文件失败:"+err.Error(), nil)
			return
		}
	}
	resourceModel.CreateAt = time.Now()
	model, err := resourceModel.Update()
	if err != nil {
		response(ctx, false, "更新资源失败:"+err.Error(), nil)
		return
	}
	resourceHistoryModel := &models.ResourceHistoryModel{
		ResourceID: resourceModel.ID,
		File:       resourceModel.File,
		Version:    resourceModel.Version,
		Path:       resourceModel.Path,
		Hash:       resourceModel.Hash,
		CreateAt:   resourceModel.CreateAt,
		Log:        logStr,
	}
	_, err = resourceHistoryModel.Insert()
	if err != nil {
		response(ctx, false, "保存历史资源版本失败:"+err.Error(), nil)
		return
	}
	response(ctx, true, "更新资源成功:", iris.Map{
		"resource": model,
	})
}

//资源列表
func ResourceLists(ctx iris.Context) {
	resourceModel := &models.ResourceModel{}
	model, err := resourceModel.All()

	if err != nil {
		response(ctx, false, "获取资源列表失败:"+err.Error(), nil)
		return
	}

	response(ctx, true, "", iris.Map{
		"resources": model,
		"timestamp": time.Now().Unix(),
	})
	return
}

//资源列表
func ResourceGroupLists(ctx iris.Context) {

	typeModel := &models.ResourceTypeModel{}
	allType, err := typeModel.All()
	if err != nil {
		response(ctx, false, "获取资源类型列表失败:"+err.Error(), nil)
		return
	}

	if len(allType) > 0 {
		resourceModel := &models.ResourceModel{}
		var data []map[string]interface{}
		for _, t := range allType {
			model, err := resourceModel.FindByType(t.ID)
			if err != nil {
				response(ctx, false, "获取资源失败:"+err.Error(), nil)
				return
			}
			if len(model) > 0 {
				resource := make(map[string]interface{})
				resource["name"] = t.Name
				resource["desc"] = t.Desc
				resource["lists"] = model
				data = append(data, resource)
			}
		}
		response(ctx, true, "", iris.Map{
			"resources": data,
			"timestamp": time.Now().Unix(),
		})
		return
	}
	response(ctx, true, "请先添加资源分类", nil)
}

//下载文件
func ResourceDownload(ctx iris.Context) {

	idStr := ctx.URLParam("id")
	version := ctx.URLParam("version")

	id, err := strconv.Atoi(idStr)
	if err != nil || version == "" {
		response(ctx, false, "文件ID和版本不能为空", nil)
		return
	}

	historyModel, err := (&models.ResourceHistoryModel{
		ResourceID: id,
		Version:    version,
	}).FirstBy()

	if err != nil {
		response(ctx, false, "当前资源不存在", nil)
		return
	}

	_, err = historyModel.Increment()
	if err != nil {
		response(ctx, false, "更新资源下载量失败", nil)
		return
	}

	userModel, _ := authUser(ctx)
	downloadModel := models.DownloadModel{
		RHId:     historyModel.ID,
		UserId:   userModel.ID,
		CreateAt: time.Now(),
	}
	_, err = downloadModel.Insert()
	if err != nil {
		response(ctx, false, "添加下载资源记录失败", nil)
		return
	}

	err = ctx.SendFile(historyModel.Path, utils.FileName(historyModel.File, historyModel.Version))
	if err != nil {
		response(ctx, false, "文件不存在"+err.Error(), nil)
	}

}
