package v1

import (
	"archive/zip"
	"fmt"
	"github.com/kataras/iris"
	"io/ioutil"
	"log"
	"os"
	"path"
	"simple-ims/models"
	"simple-ims/utils"
	"time"
)

func ProjectDelete(ctx iris.Context) {
	id := ctx.FormValue("id")
	ids := utils.StrToIntAlice(id, ",")
	if ids == nil {
		response(ctx, false, "项目ID非法", nil)
		return
	}
	pm := models.ProjectModel{}
	_, err := pm.DeleteByIds(ids)
	if err != nil {
		response(ctx, false, "删除项目失败:"+err.Error(), nil)
		return
	}

	phm := models.ProjectHistoryModel{}
	_, err = phm.DeleteByProjectId(pm.ID)
	if err != nil {
		response(ctx, false, "删除当前项目历史所有版本失败:"+err.Error(), nil)
		return
	}
	response(ctx, true, "删除项目成功", nil)
}

//项目列表
func ProjectLists(ctx iris.Context) {
	type item struct {
		models.ProjectModel
		models.ProjectHistoryModel
	}
	var row []item
	pm := models.ProjectModel{}
	model, err := pm.FindBy()
	if err != nil {
		response(ctx, true, "获取项目列表失败:"+err.Error(), nil)
		return
	}
	for _, v := range *model {
		phm := models.ProjectHistoryModel{
			ProjectId: v.ID,
		}
		//todo
		first, err := phm.First()
		if err != nil {
			response(ctx, true, "获取项目版本失败:"+err.Error(), nil)
			return
		}
		row = append(row, item{v, *first})
	}
	response(ctx, true, "获取项目列表成功", row)
}

//添加项目版本
func ProjectUpgrade(ctx iris.Context) {
	projectId, err := ctx.PostValueInt("project_id")
	if err != nil {
		response(ctx, false, "项目ID非法", nil)
		return
	}
	version := ctx.FormValue("version")
	RHIds := ctx.FormValue("rh_ids")
	logStr := ctx.FormValue("log")
	if version == "" || RHIds == "" || logStr == "" {
		response(ctx, false, "请输入版本号和日志,选择对应资源", nil)
		return
	}
	phm := models.ProjectHistoryModel{
		ProjectId: projectId,
		Version:   version,
		Log:       logStr,
		RHIds:     RHIds,
		Hash:      utils.Md5Str(string(projectId) + version + logStr + RHIds),
	}
	model, err := phm.Insert()
	if err != nil {
		response(ctx, false, "保存项目版本失败:"+err.Error(), nil)
		return
	}
	//var sourcePath []string
	pm := models.ProjectModel{
		ID: projectId,
	}
	projectModel, err := pm.FirstBy()
	if err != nil {
		log.Println(err)
		return
	}
	uploadDir := "uploads/" + time.Now().Format("2006/01/")
	if !utils.Mkdir(uploadDir) {
		response(ctx, false, "创建文件夹失败", nil)
		return
	}

	fZip, err := os.Create(uploadDir + utils.FileName(projectModel.Name, model.Version) + ".zip")
	if err != nil {
		log.Println("Create", err)
		return
	}
	w := zip.NewWriter(fZip)
	defer w.Close()
	for _, id := range utils.StrToIntAlice(model.RHIds, ",") {
		log.Println("id", id)
		rhm := models.ResourceHistoryModel{
			ID: id,
		}
		resourceHistoryModel, err := rhm.FirstBy()
		if err != nil {
			log.Println("FirstBy", err)
			return
		}
		log.Println(resourceHistoryModel.Path)
		fw, err := w.Create(path.Base(resourceHistoryModel.Path))
		if err != nil {
			log.Println("Create", err)
			return
		}
		fileContent, err := ioutil.ReadFile(resourceHistoryModel.Path)
		fmt.Println(string(fileContent))
		if err != nil {
			log.Println("ReadFile", err)
			return
		}
		_, err = fw.Write(fileContent)
		if err != nil {
			log.Println("Write", err)
			return
		}
	}

	pm.PHId = model.ID
	_, err = pm.Update()
	if err != nil {
		response(ctx, false, "更新项目失败:"+err.Error(), model)
		return
	}
	response(ctx, true, "保存项目版本成功", model)
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

func ProjectDownload(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	version := ctx.URLParam("version")
	if err != nil || version == "" {
		response(ctx, false, "文件ID和版本不能为空", nil)
		return
	}
	phm := models.ProjectHistoryModel{
		ProjectId: id,
		Version:   version,
	}

	model, err := phm.First()
	if err != nil {
		response(ctx, false, "当前项目版本不存在", nil)
		return
	}
	RHIds := utils.StrToIntAlice(model.RHIds, ",")
	_ = RHIds
	//for _, id := range RHIds {
	//	rhm := models.ResourceHistoryModel{
	//		ID: id,
	//	}
	//	resourceHistoryModel, err := rhm.FirstBy()
	//	if err != nil {
	//
	//		return
	//	}
	//
	//}
	//rhm := models.ResourceHistoryModel{
	//	ResourceID: id,
	//	Version:    version,
	//}
	//a, err := rhm.FirstBy()
	//if err != nil {
	//	response(ctx, false, "当前资源不存在", nil)
	//	return
	//}
	//
	//_, err = historyModel.Increment()
	//if err != nil {
	//	response(ctx, false, "更新资源下载量失败", nil)
	//	return
	//}
}
