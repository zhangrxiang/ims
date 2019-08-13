package v1

import (
	"github.com/kataras/iris"
	"log"
	"os"
	"path"
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
		resourceModel.Path = uploadDir + resourceModel.Hash + path.Ext(info.Filename)
		resourceModel.CreateAt = time.Now()
		err = utils.CopyFile(resourceModel.Path, file)
		if err != nil {
			response(ctx, false, "保存文件失败:"+err.Error(), nil)
			return
		}
	}

	model, err := resourceModel.Insert()

	if err != nil {
		response(ctx, false, "添加资源失败:", nil)
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
	_id := ctx.FormValue("id")
	name := ctx.FormValue("name")
	_type := ctx.FormValue("type")
	version := ctx.FormValue("version")
	desc := ctx.FormValue("desc")

	if name == "" || _type == "" {
		response(ctx, false, "请输入资源名称,选择资源类型", nil)
		return
	}

	id, err := strconv.Atoi(_id)
	if err != nil {
		response(ctx, false, "资源ID格式不合法:"+err.Error(), nil)
		return
	}
	t, err := strconv.Atoi(_type)
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
		resourceModel.Path = uploadDir + resourceModel.Hash + path.Ext(info.Filename)
		resourceModel.CreateAt = time.Now()
		err = utils.CopyFile(resourceModel.Path, file)
		if err != nil {
			response(ctx, false, "保存文件失败:"+err.Error(), nil)
			return
		}
	}

	model, err := resourceModel.Update()

	if err != nil {
		response(ctx, false, "添加资源失败:"+err.Error(), nil)
		return
	}
	response(ctx, true, "保存文件成功:", iris.Map{
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

	if err == nil && allType != nil {
		resourceModel := &models.ResourceModel{}
		var data []map[string]interface{}
		for _, t := range allType {
			model, err := resourceModel.FindByType(t.ID)
			if model != nil && err == nil {
				resource := make(map[string]interface{})
				resource["name"] = t.Name
				resource["desc"] = t.Desc
				resource["lists"] = model
				data = append(data, resource)
			}
			if err == models.NoRecordExists {
				response(ctx, true, "无可用资源", nil)
				return
			}

			if err != nil {
				response(ctx, false, "获取资源失败:"+err.Error(), nil)
				return
			}
		}
		response(ctx, true, "", iris.Map{
			"resources": data,
			"timestamp": time.Now().Unix(),
		})
		return
	}
	if err == models.NoRecordExists {
		response(ctx, true, "请先添加资源分类", nil)
		return
	}
	if err != nil {
		response(ctx, false, "获取资源类型列表失败:"+err.Error(), nil)
		return
	}

}

//下载文件
func ResourceDownload(ctx iris.Context) {
	p := ctx.URLParam("path")
	f := ctx.URLParam("file")

	if p == "" || f == "" {
		response(ctx, false, "文件路径不能为空", nil)
		return
	}

	err := ctx.SendFile(p, f)

	if err != nil {
		response(ctx, false, "文件不存在", nil)
	}

}
